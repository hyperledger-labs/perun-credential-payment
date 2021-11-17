package connection

import (
	"bytes"
	"context"
	"fmt"

	"github.com/perun-network/verifiable-credential-payment/app"
	"perun.network/go-perun/channel"
)

type CredentialRequest struct {
	resp    chan CredentialRequestResponse
	DocHash app.Hash
	Price   channel.Bal
}

func (r *CredentialRequest) CheckDoc(doc []byte) error {
	docHash := app.ComputeDocumentHash(doc)
	if !bytes.Equal(docHash[:], r.DocHash[:]) {
		return fmt.Errorf("wrong document")
	}
	return nil
}

func (r *CredentialRequest) Accept(ctx context.Context) error {
	errs := make(chan error)
	r.resp <- &CredentialRequestResponseAccept{ctx, errs}
	return <-errs
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
