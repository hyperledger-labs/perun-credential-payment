package perun

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"perun.network/go-perun/backend/ethereum/channel"
	"perun.network/go-perun/backend/ethereum/wallet"
	wtest "perun.network/go-perun/backend/ethereum/wallet/simple"
	"perun.network/go-perun/client"
	"perun.network/go-perun/watcher/local"
	"perun.network/go-perun/wire"
	"perun.network/go-perun/wire/net"
	"perun.network/go-perun/wire/net/simple"
)

type Peer struct {
	Peer    wire.Address
	Address string
}

type ClientConfig struct {
	PrivateKey    *ecdsa.PrivateKey
	Host          string
	ETHNodeURL    string
	Adjudicator   common.Address
	AssetHolder   common.Address
	DialerTimeout time.Duration
	Peers         []Peer
	TxFinality    uint64
	ChainID       *big.Int
}

type Client struct {
	EthClient       *ethclient.Client
	PerunClient     *client.Client
	Bus             *net.Bus
	Listener        net.Listener
	ContractBackend channel.ContractInterface
	Wallet          *wtest.Wallet
	Account         *wtest.Account
}

func SetupClient(ctx context.Context, cfg ClientConfig) (*Client, error) {
	// Create wallet and account
	w := wtest.NewWallet(cfg.PrivateKey)
	addr := wallet.AsWalletAddr(crypto.PubkeyToAddress(cfg.PrivateKey.PublicKey))
	pAccount, err := w.Unlock(addr)
	if err != nil {
		panic("failed to create account")
	}
	account := pAccount.(*wtest.Account)

	// Create Ethereum client and contract backend
	ethClient, cb, err := createContractBackend(cfg.ETHNodeURL, w, cfg.ChainID, cfg.TxFinality)
	if err != nil {
		return nil, errors.WithMessage(err, "creating contract backend")
	}

	// Setup adjudicator.
	if err := channel.ValidateAdjudicator(ctx, cb, cfg.Adjudicator); err != nil {
		return nil, fmt.Errorf("validating adjudicator: %w", err)
	}
	adjudicator := channel.NewAdjudicator(cb, cfg.Adjudicator, account.Account.Address, account.Account)

	// Setup asset holder.
	funder := createFunder(cb, account.Account, cfg.AssetHolder)

	// Setup network.
	listener, bus, err := setupNetwork(account, cfg.Host, cfg.Peers, cfg.DialerTimeout)
	if err != nil {
		return nil, errors.WithMessage(err, "setting up network")
	}

	// Setup watcher.
	watcher, err := local.NewWatcher(adjudicator)
	if err != nil {
		return nil, fmt.Errorf("initializing watcher: %w", err)
	}

	// Initialize Perun client.
	c, err := client.New(account.Address(), bus, funder, adjudicator, w, watcher)
	if err != nil {
		return nil, errors.WithMessage(err, "initializing client")
	}

	return &Client{ethClient, c, bus, listener, cb, w, account}, nil
}

func createContractBackend(nodeURL string, wallet *wtest.Wallet, chainID *big.Int, txFinality uint64) (*ethclient.Client, channel.ContractBackend, error) {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		return nil, channel.ContractBackend{}, nil
	}

	signer := types.NewEIP155Signer(chainID)
	tr := wtest.NewTransactor(wallet, signer)

	return client, channel.NewContractBackend(client, tr, txFinality), nil
}

func setupNetwork(account wire.Account, host string, peerAddresses []Peer, dialerTimeout time.Duration) (listener net.Listener, bus *net.Bus, err error) {
	dialer := simple.NewTCPDialer(dialerTimeout)

	for _, pa := range peerAddresses {
		dialer.Register(pa.Peer, pa.Address)
	}

	listener, err = simple.NewTCPListener(host)
	if err != nil {
		err = fmt.Errorf("creating listener: %w", err)
		return
	}

	bus = net.NewBus(account, dialer)
	return listener, bus, nil
}

func createFunder(cb channel.ContractBackend, account accounts.Account, assetHolder common.Address) *channel.Funder {
	f := channel.NewFunder(cb)
	asset := wallet.Address(assetHolder)
	depositor := new(channel.ETHDepositor)
	f.RegisterAsset(asset, depositor, account)
	return f
}
