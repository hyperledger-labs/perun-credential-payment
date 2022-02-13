package channel

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/perun-network/perun-credential-payment/app"
	"github.com/perun-network/perun-credential-payment/app/data"
	"perun.network/go-perun/backend/ethereum/wallet/simple"
	"perun.network/go-perun/client"
)

type CredentialRequest struct {
	resp    chan CredentialRequestResponse
	offer   *data.Offer
	channel *Channel
}

func (r *CredentialRequest) CheckDoc(doc []byte) error {
	docHash := app.ComputeDocumentHash(doc)
	if !bytes.Equal(docHash[:], r.offer.DataHash[:]) {
		return fmt.Errorf("wrong document")
	}
	return nil
}

func (r *CredentialRequest) CheckPrice(p *big.Int) error {
	if r.offer.Price.Cmp(p) != 0 {
		return fmt.Errorf("wrong price")
	}
	return nil
}

func (r *CredentialRequest) IssueCredential(ctx context.Context, acc *simple.Account) error {
	errs := make(chan error)
	r.resp <- &CredentialRequestResponseAccept{ctx, errs}
	err := <-errs
	if err != nil {
		return fmt.Errorf("accepting credential request: %w", err)
	}

	// Issue credential.
	err = r.channel.issueCredential(ctx, r.offer, acc)
	if err != nil {
		return fmt.Errorf("issueing credential: %w", err)
	}

	return nil
}

type (
	CredentialRequestResponse interface {
		Context() context.Context
		Result() chan error
	}

	CredentialRequestResponseAccept struct {
		ctx  context.Context
		errs chan error
	}
)

func (r *CredentialRequestResponseAccept) Context() context.Context {
	return r.ctx
}

func (r *CredentialRequestResponseAccept) Result() chan error {
	return r.errs
}

type AsyncCredential struct {
	sigRegCallback
}

func (c *AsyncCredential) Await(ctx context.Context) (*CredentialProposal, error) {
	select {
	case prop := <-c.sigRegCallback:
		return prop, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type CredentialProposal struct {
	*client.UpdateResponder
	Signature []byte
}
