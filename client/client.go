package client

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	pkgapp "github.com/perun-network/verifiable-credential-payment/app"
	"github.com/perun-network/verifiable-credential-payment/client/connection"
	"github.com/perun-network/verifiable-credential-payment/client/perun"
	"github.com/pkg/errors"
	"perun.network/go-perun/backend/ethereum/bindings/assetholdereth"
	ethchannel "perun.network/go-perun/backend/ethereum/channel"
	ethwallet "perun.network/go-perun/backend/ethereum/wallet"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/client"
	"perun.network/go-perun/wire"
)

type ClientConfig struct {
	perun.ClientConfig
	ChallengeDuration time.Duration
	AppAddress        common.Address
	Honest            bool
}

type PaymentAcceptancePolicy = func(
	amount *big.Int,
	collateral *big.Int,
	funding *big.Int,
	balance *big.Int,
	hasOverdrawn bool,
) (ok bool)

type Client struct {
	perunClient       *perun.Client
	assetHolderAddr   common.Address
	assetHolder       *assetholdereth.AssetHolderETH
	challengeDuration time.Duration
	appAddress        common.Address
	channelProposals  chan *connection.ChannelProposal
	connections       *connection.Registry
	honest            bool
}

func StartClient(ctx context.Context, cfg ClientConfig) (*Client, error) {
	perunClient, err := perun.SetupClient(ctx, cfg.ClientConfig)
	if err != nil {
		return nil, errors.WithMessage(err, "creating perun client")
	}

	if err := ethchannel.ValidateAssetHolderETH(ctx, perunClient.ContractBackend, cfg.AssetHolder, cfg.Adjudicator); err != nil {
		return nil, fmt.Errorf("validating adjudicator: %w", err)
	}
	ah, err := assetholdereth.NewAssetHolderETH(cfg.AssetHolder, perunClient.ContractBackend)
	if err != nil {
		return nil, errors.WithMessage(err, "loading asset holder")
	}

	c := &Client{
		perunClient:       perunClient,
		assetHolderAddr:   cfg.AssetHolder,
		assetHolder:       ah,
		challengeDuration: cfg.ChallengeDuration,
		appAddress:        cfg.AppAddress,
		channelProposals:  make(chan *connection.ChannelProposal),
		connections:       connection.NewRegistry(),
		honest:            cfg.Honest,
	}

	h := &handler{Client: c}

	go c.perunClient.PerunClient.Handle(h, h)
	go c.perunClient.Bus.Listen(c.perunClient.Listener)

	return c, nil
}

func (c *Client) Connect(ctx context.Context, peer wire.Address, balance channel.Bal) (*connection.Connection, error) {
	app := pkgapp.NewCredentialSwapApp(ethwallet.AsWalletAddr(c.appAddress))
	peers := []wire.Address{c.perunClient.Account.Address(), peer}
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

	ch, err := c.perunClient.PerunClient.ProposeChannel(ctx, prop)
	if err != nil {
		return nil, fmt.Errorf("proposing channel: %w", err)
	}
	conn := connection.NewConnection(ch)
	c.connections.Add(conn)

	h := connection.NewEventHandler(conn)
	go func() {
		err := conn.Watch(h)
		if err != nil {
			c.Logf("Watching failed: %v", err)
		}
	}()

	return conn, nil
}

func (c *Client) NextConnectionRequest(ctx context.Context) (*connection.ConnectionRequest, error) {
	p, ok := <-c.channelProposals
	if !ok {
		return nil, fmt.Errorf("channel closed")
	}
	return connection.NewConnectionRequest(p, c.PerunAddress(), c.connections), nil
}

func (c *Client) Honest() bool {
	return c.honest
}

func (c *Client) Shutdown() {
	c.perunClient.PerunClient.Close()
	c.perunClient.Bus.Close()
}
