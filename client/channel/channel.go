package channel

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/perun-network/perun-credential-payment/app"
	"github.com/perun-network/perun-credential-payment/app/data"
	"github.com/perun-network/perun-credential-payment/pkg/atomic"
	ethwallet "perun.network/go-perun/backend/ethereum/wallet"
	swallet "perun.network/go-perun/backend/ethereum/wallet/simple"
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

type ChannelRequest struct {
	p        *ChannelProposal
	acc      wallet.Address
	registry *Registry
}

func NewChannelRequest(
	p *ChannelProposal,
	acc wallet.Address,
	registry *Registry,
) *ChannelRequest {
	return &ChannelRequest{
		p:        p,
		acc:      acc,
		registry: registry,
	}
}

func (r *ChannelRequest) Peer() wallet.Address {
	return r.p.p.Participant
}

func (r *ChannelRequest) Accept(ctx context.Context) (*Channel, error) {
	msg := r.p.p.Accept(r.acc, client.WithRandomNonce())
	perunCh, err := r.p.r.Accept(ctx, msg)
	if err != nil {
		return nil, fmt.Errorf("accepting channel: %w", err)
	}
	ch := NewChannel(perunCh)
	r.registry.Add(ch)

	h := NewEventHandler(ch)
	go func() {
		err := ch.Watch(h)
		if err != nil {
			ch.Log().Warnf("Watching failed: %v", err)
		}
	}()

	return ch, nil
}

type Channel struct {
	*client.Channel
	sigs         *sigReg
	credRequests chan *CredentialRequest
	disputed     *atomic.Bool
	concludable  *atomic.Bool
	concluded    *atomic.Bool
}

func NewChannel(ch *client.Channel) *Channel {
	return &Channel{
		Channel:      ch,
		sigs:         newSigReg(),
		credRequests: make(chan *CredentialRequest),
		disputed:     atomic.NewBool(false),
		concludable:  atomic.NewBool(false),
		concluded:    atomic.NewBool(false),
	}
}

func (ch *Channel) Disputed() bool {
	return ch.disputed.Value()
}

func (ch *Channel) RequestCredential(
	ctx context.Context,
	doc []byte,
	price channel.Bal,
	issuer *ethwallet.Address,
) (*AsyncCredential, error) {
	// Compute hash.
	h := app.ComputeDocumentHash(doc)

	issuerEthAddr := common.Address(*issuer)
	callback, err := ch.sigs.RegisterCallback(h, issuerEthAddr)
	if err != nil {
		return nil, err
	}

	// Perform request.
	err = ch.UpdateBy(ctx, func(s *channel.State) error {
		s.Data = &data.Offer{
			Issuer:   issuerEthAddr,
			DataHash: h,
			Price:    price,
			Buyer:    uint16(ch.Idx()),
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("updating channel: %w", err)
	}

	return &AsyncCredential{callback}, nil
}

func (ch *Channel) addCredentialRequest(offer *data.Offer) chan CredentialRequestResponse {
	response := make(chan CredentialRequestResponse)
	ch.credRequests <- &CredentialRequest{
		resp:    response,
		offer:   offer,
		channel: ch,
	}
	return response
}

func (ch *Channel) NextCredentialRequest(ctx context.Context) (*CredentialRequest, error) {
	select {
	case r := <-ch.credRequests:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (ch *Channel) addSignature(sig app.Signature, h app.Hash, issuer common.Address, responder *client.UpdateResponder) {
	ch.sigs.Push(sig, h, issuer, responder)
}

func (ch *Channel) issueCredential(ctx context.Context, offer *data.Offer, acc *swallet.Account) error {
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
		s.Allocation.AddToBalance(ch.Idx(), asset, offer.Price)

		return nil
	}

	err := ch.UpdateBy(ctx, up)
	if err != nil {
		ch.Log().Warnf("Failed to update channel off-ledger: %v", err)
		ch.Log().Warnf("Forcing update on-ledger")

		ch.disputed.SetValue(true)
		err := ch.ForceUpdate(ctx, func(s *channel.State) {
			err := up(s)
			if err != nil {
				ch.Log().Warnf("Updating channel state: %v", err)
			}
		})
		if err != nil {
			return fmt.Errorf("forcing update: %w", err)
		}
	}

	return nil
}

func (ch *Channel) TryClose(ctx context.Context, attempts int) error {
	for i := 1; i <= attempts; i++ {
		err := ch.Close(ctx)
		if err == nil {
			return nil
		} else {
			ch.Log().Warnf("Failed to close channel (attempt %d): %v", i, err)
		}
	}
	return fmt.Errorf("Failed to close channel in %d attempts", attempts)
}

func (ch *Channel) Close(ctx context.Context) error {
	if ch.Disputed() {
		// If there is a dispute, we wait until the channel is concludable.
		err := ch.WaitConcludadable(ctx)
		if err != nil {
			return fmt.Errorf("waiting for channel concludable: %w", err)
		}
	} else if !ch.State().IsFinal {
		// If there is no dispute, we attempt to finalize the channel.
		err := ch.UpdateBy(ctx, func(s *channel.State) error {
			s.Data = &data.DefaultData{}
			s.IsFinal = true
			return nil
		})
		if err != nil {
			ch.Log().Warnf("Failed to finalize channel off-ledger: %v", err)
		}
	}

	err := ch.Settle(ctx, false)
	if err != nil {
		return fmt.Errorf("settling: %w", err)
	}

	return nil
}

func (ch *Channel) WaitConcludadable(ctx context.Context) error {
	return waitCondition(ctx, func() bool {
		return ch.State().IsFinal || ch.concludable.Value()
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
