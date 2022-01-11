package connection

import (
	"context"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/perun-network/perun-credential-payment/app"
	"perun.network/go-perun/client"
)

type (
	sigRegKey struct {
		Issuer  common.Address
		DocHash app.Hash
	}

	sigRegReturnVal = *CredentialProposal

	sigReg struct {
		sync.RWMutex
		callbacks map[sigRegKey]chan sigRegReturnVal
	}
)

func newSigReg() *sigReg {
	return &sigReg{
		callbacks: make(map[sigRegKey]chan sigRegReturnVal),
	}
}

func (r *sigReg) RegisterCallback(h app.Hash, issuer common.Address) (sigRegCallback, error) {
	r.Lock()
	defer r.Unlock()

	k := sigRegKey{Issuer: issuer, DocHash: h}
	callback := make(chan sigRegReturnVal, 1)

	_, ok := r.callbacks[k]
	if ok {
		return nil, fmt.Errorf("already registered")
	}

	r.callbacks[k] = callback
	return sigRegCallback(callback), nil
}

func (r *sigReg) Push(sig app.Signature, h app.Hash, issuer common.Address, responder *client.UpdateResponder) {
	r.Lock()
	defer r.Unlock()

	k := sigRegKey{Issuer: issuer, DocHash: h}
	cb, ok := r.callbacks[k]
	if !ok {
		return
	}

	cb <- &CredentialProposal{
		UpdateResponder: responder,
		Signature:       sig,
	}
	delete(r.callbacks, k)
}

type sigRegCallback chan sigRegReturnVal

func (cb sigRegCallback) Await(ctx context.Context) (sigRegReturnVal, error) {
	select {
	case r := <-cb:
		return r, nil
	case <-ctx.Done():
		return nil, fmt.Errorf("waiting for callback: %w", ctx.Err())
	}
}
