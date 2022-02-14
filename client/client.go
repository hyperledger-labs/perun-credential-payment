package client

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	pkgapp "github.com/perun-network/perun-credential-payment/app"
	clientchannel "github.com/perun-network/perun-credential-payment/client/channel"
	"github.com/pkg/errors"
	ethchannel "perun.network/go-perun/backend/ethereum/channel"
	ethwallet "perun.network/go-perun/backend/ethereum/wallet"
	"perun.network/go-perun/backend/ethereum/wallet/simple"
	wtest "perun.network/go-perun/backend/ethereum/wallet/simple"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/client"
	"perun.network/go-perun/wallet"
	"perun.network/go-perun/watcher/local"
	"perun.network/go-perun/wire"
)

type ClientConfig struct {
	PrivateKey        *ecdsa.PrivateKey
	ETHNodeURL        string
	Adjudicator       common.Address
	AssetHolder       common.Address
	AppAddress        common.Address
	TxFinality        uint64
	ChainID           *big.Int
	Bus               *wire.LocalBus
	ChallengeDuration time.Duration
}

type Client struct {
	perunClient       *client.Client
	assetHolderAddr   common.Address
	challengeDuration time.Duration
	appAddress        common.Address
	channelProposals  chan *clientchannel.ChannelProposal
	channels          *clientchannel.Registry
	account           *wtest.Account
}

// StartClient create a new channel client and starts it up.
func StartClient(ctx context.Context, cfg ClientConfig) (*Client, error) {
	// Create wallet and account.
	w := wtest.NewWallet(cfg.PrivateKey)
	addr := ethwallet.AsWalletAddr(crypto.PubkeyToAddress(cfg.PrivateKey.PublicKey))
	pAccount, err := w.Unlock(addr)
	if err != nil {
		panic("failed to create account")
	}
	account := pAccount.(*wtest.Account)

	// Create contract backend.
	ethClient, err := ethclient.Dial(cfg.ETHNodeURL)
	if err != nil {
		return nil, fmt.Errorf("dialing Ethereum node: %w", err)
	}
	signer := types.NewEIP155Signer(cfg.ChainID)
	tr := wtest.NewTransactor(w, signer)
	cb := ethchannel.NewContractBackend(ethClient, tr, cfg.TxFinality)

	// Setup adjudicator.
	if err := ethchannel.ValidateAdjudicator(ctx, cb, cfg.Adjudicator); err != nil {
		return nil, fmt.Errorf("validating adjudicator: %w", err)
	}
	adjudicator := ethchannel.NewAdjudicator(cb, cfg.Adjudicator, account.Account.Address, account.Account)

	// Setup funder.
	if err := ethchannel.ValidateAssetHolderETH(ctx, cb, cfg.AssetHolder, cfg.Adjudicator); err != nil {
		return nil, fmt.Errorf("validating adjudicator: %w", err)
	}
	funder := ethchannel.NewFunder(cb)
	asset, depositor := ethwallet.Address(cfg.AssetHolder), new(ethchannel.ETHDepositor)
	funder.RegisterAsset(asset, depositor, account.Account)

	// Setup watcher.
	watcher, err := local.NewWatcher(adjudicator)
	if err != nil {
		return nil, fmt.Errorf("initializing watcher: %w", err)
	}

	// Initialize Perun client.
	perunClient, err := client.New(account.Address(), cfg.Bus, funder, adjudicator, w, watcher)
	if err != nil {
		return nil, errors.WithMessage(err, "initializing client")
	}

	c := &Client{
		perunClient:       perunClient,
		assetHolderAddr:   cfg.AssetHolder,
		challengeDuration: cfg.ChallengeDuration,
		appAddress:        cfg.AppAddress,
		channelProposals:  make(chan *clientchannel.ChannelProposal),
		channels:          clientchannel.NewRegistry(),
		account:           account,
	}

	// Start request handler.
	h := &handler{Client: c}
	go c.perunClient.Handle(h, h)

	return c, nil
}

func (c *Client) OpenChannel(ctx context.Context, peer wire.Address, balance channel.Bal) (*clientchannel.Channel, error) {
	app := pkgapp.NewCredentialSwapApp(ethwallet.AsWalletAddr(c.appAddress))
	peers := []wire.Address{c.account.Address(), peer}
	withApp := client.WithApp(app, app.InitData())

	asset := ethwallet.AsWalletAddr(c.assetHolderAddr)
	alloc := channel.NewAllocation(2, asset)
	ourIndex, peerIndex := channel.Index(0), channel.Index(1)
	alloc.SetBalance(ourIndex, asset, balance)
	alloc.SetBalance(peerIndex, asset, big.NewInt(0))

	prop, err := client.NewLedgerChannelProposal(
		c.challengeDurationInSeconds(),
		c.PerunAddress(),
		alloc,
		peers,
		withApp,
	)
	if err != nil {
		return nil, fmt.Errorf("creating channel proposal: %w", err)
	}

	perunCh, err := c.perunClient.ProposeChannel(ctx, prop)
	if err != nil {
		return nil, fmt.Errorf("proposing channel: %w", err)
	}
	ch := clientchannel.NewChannel(perunCh)
	c.channels.Add(ch)

	h := clientchannel.NewEventHandler(ch)
	go func() {
		err := ch.Watch(h)
		if err != nil {
			c.Logf("Watching failed: %v", err)
		}
	}()

	return ch, nil
}

func (c *Client) NextChannelRequest(ctx context.Context) (*clientchannel.ChannelRequest, error) {
	p, ok := <-c.channelProposals
	if !ok {
		return nil, fmt.Errorf("channel closed")
	}
	return clientchannel.NewChannelRequest(p, c.PerunAddress(), c.channels), nil
}

func (c *Client) Shutdown() {
	c.perunClient.Close()
}

func (c *Client) Account() *simple.Account {
	return c.account
}

func (c *Client) PerunAddress() wallet.Address {
	return c.account.Address()
}

func (c *Client) EthAddress() common.Address {
	return c.account.Account.Address
}

func (c *Client) challengeDurationInSeconds() uint64 {
	return uint64(c.challengeDuration.Seconds())
}

func (c *Client) Logf(format string, v ...interface{}) {
	log.Printf("Client %v: %v", c.EthAddress(), fmt.Sprintf(format, v...))
}
