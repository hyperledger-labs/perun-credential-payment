package connection

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/perun-network/verifiable-credential-payment/app"
	"github.com/perun-network/verifiable-credential-payment/app/data"
	"github.com/perun-network/verifiable-credential-payment/pkg/atomic"
	ewallet "perun.network/go-perun/backend/ethereum/wallet/simple"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/client"
	"perun.network/go-perun/wallet"
)

type ChannelProposal struct {
	p *client.LedgerChannelProposal
	r *client.ProposalResponder
}

func NewChannelProposal(
	p *client.LedgerChannelProposal,
	r *client.ProposalResponder,
) *ChannelProposal {
	return &ChannelProposal{
		p: p,
		r: r,
	}
}

type ConnectionRequest struct {
	p        *ChannelProposal
	acc      wallet.Address
	registry *Registry
}

func NewConnectionRequest(
	p *ChannelProposal,
	acc wallet.Address,
	registry *Registry,
) *ConnectionRequest {
	return &ConnectionRequest{
		p:        p,
		acc:      acc,
		registry: registry,
	}
}

func (r *ConnectionRequest) Peer() wallet.Address {
	return r.p.p.Participant
}

func (r *ConnectionRequest) Accept(ctx context.Context) (*Connection, error) {
	msg := r.p.p.Accept(r.acc, client.WithRandomNonce())
	ch, err := r.p.r.Accept(ctx, msg)
	if err != nil {
		return nil, fmt.Errorf("accepting channel: %w", err)
	}
	conn := NewConnection(ch)
	r.registry.Add(conn)

	h := NewEventHandler(conn)
	go func() {
		err := conn.Watch(h)
		if err != nil {
			conn.Log().Warnf("Watching failed: %v", err)
		}
	}()

	return conn, nil
}

type Connection struct {
	*client.Channel
	sigs         *sigReg
	credRequests chan *CredentialRequest
	disputed     *atomic.Bool
	concludable  *atomic.Bool
	concluded    *atomic.Bool
}

func NewConnection(ch *client.Channel) *Connection {
	return &Connection{
		Channel:      ch,
		sigs:         newSigReg(),
		credRequests: make(chan *CredentialRequest),
		disputed:     atomic.NewBool(false),
		concludable:  atomic.NewBool(false),
		concluded:    atomic.NewBool(false),
	}
}

func (c *Connection) Disputed() bool {
	return c.disputed.Value()
}

func (c *Connection) RequestCredential(
	ctx context.Context,
	doc []byte,
	price channel.Bal,
	issuer common.Address,
) (*AsyncCredential, error) {
	// Compute hash.
	h := app.ComputeDocumentHash(doc)

	callback, err := c.sigs.RegisterCallback(h, issuer)
	if err != nil {
		return nil, err
	}

	// Perform request.
	err = c.UpdateBy(ctx, func(s *channel.State) error {
		s.Data = &data.Offer{
			Issuer:   issuer,
			DataHash: h,
			Price:    price,
			Buyer:    uint16(c.Idx()),
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("updating channel: %w", err)
	}

	return &AsyncCredential{callback}, nil
}

func (c *Connection) addCredentialRequest(offer *data.Offer) chan CredentialRequestResponse {
	response := make(chan CredentialRequestResponse)
	c.credRequests <- &CredentialRequest{
		resp:  response,
		offer: offer,
		conn:  c,
	}
	return response
}

func (c *Connection) NextCredentialRequest(ctx context.Context) (*CredentialRequest, error) {
	select {
	case r := <-c.credRequests:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (c *Connection) addSignature(sig app.Signature, h app.Hash, issuer common.Address, responder *client.UpdateResponder) {
	c.sigs.Push(sig, h, issuer, responder)
}

func (c *Connection) issueCredential(ctx context.Context, offer *data.Offer, acc *ewallet.Account) error {
	up := func(s *channel.State) error {
		// Check inputs against current state.
		curOffer, ok := s.Data.(*data.Offer)
		if !ok {
			return fmt.Errorf("data has wrong type: %T", s.Data)
		} else if !curOffer.Equal(offer) {
			return fmt.Errorf("unequal offers: got %v, expected %v", curOffer, offer)
		} else if addr := acc.Account.Address; offer.Issuer != addr {
			return fmt.Errorf("unequal addresses: got %v, expected %v", addr, offer.Issuer)
		}

		// Sign.
		sig, err := app.SignHash(acc, offer.DataHash)
		if err != nil {
			return fmt.Errorf("signing hash: %w", err)
		}

		// Update state data.
		var cert data.Cert
		copy(cert.Signature[:], sig[:])
		s.Data = &cert

		// Update balances.
		asset := s.Allocation.Assets[app.AssetIdx]
		s.Allocation.SubFromBalance(channel.Index(offer.Buyer), asset, offer.Price)
		s.Allocation.AddToBalance(c.Idx(), asset, offer.Price)

		return nil
	}

	err := c.UpdateBy(ctx, up)
	if err != nil {
		c.Log().Warnf("Failed to update channel off-ledger: %v", err)
		c.Log().Warnf("Forcing update on-ledger")

		c.disputed.SetValue(true)
		err := c.ForceUpdate(ctx, func(s *channel.State) {
			err := up(s)
			if err != nil {
				c.Log().Warnf("Updating channel state: %v", err)
			}
		})
		if err != nil {
			return fmt.Errorf("forcing update: %w", err)
		}
	}

	return nil
}

func (c *Connection) TryClose(ctx context.Context, attempts int) error {
	for i := 1; i <= attempts; i++ {
		err := c.Close(ctx)
		if err == nil {
			return nil
		} else {
			c.Log().Warnf("Failed to close channel (attempt %d): %v", i, err)
		}
	}
	return fmt.Errorf("Failed to close channel in %d attempts", attempts)
}

func (c *Connection) Close(ctx context.Context) error {
	if c.Disputed() {
		// If there is a dispute, we wait until the channel is concludable.
		err := c.WaitConcludadable(ctx)
		if err != nil {
			return fmt.Errorf("waiting for channel concludable: %w", err)
		}
	} else if !c.State().IsFinal {
		// If there is no dispute, we attempt to finalize the channel.
		err := c.UpdateBy(ctx, func(s *channel.State) error {
			s.Data = &data.DefaultData{}
			s.IsFinal = true
			return nil
		})
		if err != nil {
			c.Log().Warnf("Failed to finalize channel off-ledger: %v", err)
		}
	}

	err := c.Settle(ctx, false)
	if err != nil {
		return fmt.Errorf("settling: %w", err)
	}

	return nil
}

func (c *Connection) WaitConcludadable(ctx context.Context) error {
	return waitCondition(ctx, func() bool {
		return c.State().IsFinal || c.concludable.Value()
	})
}

func waitCondition(ctx context.Context, cond func() bool) error {
	const tick = 500 * time.Millisecond

loop:
	for {
		select {
		case <-time.After(tick):
			if cond() {
				break loop
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}
