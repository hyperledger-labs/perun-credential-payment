package main_test

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"testing"

	"github.com/perun-network/verifiable-credential-payment/client/connection"
	"github.com/perun-network/verifiable-credential-payment/test"
	"github.com/stretchr/testify/require"
)

func TestCredentialSwap(t *testing.T) {
	t.Run("Honest holder", func(t *testing.T) {
		runCredentialSwapTest(t, true)
	})
	t.Run("Dishonest holder", func(t *testing.T) {
		runCredentialSwapTest(t, false)
	})
}

func runCredentialSwapTest(t *testing.T, honestHolder bool) {
	require := require.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	// Setup test environment.
	env := test.Setup(t, honestHolder)
	env.LogAccountBalances()
	wg, errs := sync.WaitGroup{}, make(chan error)
	wg.Add(2)
	holder, issuer := env.Holder, env.Issuer

	doc := []byte("Perun/Bosch: SSI Credential Payment")
	balance := test.EthToWei(big.NewFloat(5))
	price := test.EthToWei(big.NewFloat(1))

	// Run credential holder.
	go func() {
		conn, err := holder.Connect(ctx, issuer.PerunAddress(), balance)
		if err != nil {
			errs <- fmt.Errorf("proposing connection: %w", err)
			return
		}

		cred, err := conn.BuyCredential(ctx, doc, price, issuer.Address())
		if err != nil {
			errs <- fmt.Errorf("buying credential: %w", err)
			return
		}

		holder.Logf("obtained credential: %v", cred.String())

		// If dishonest, wait for dispute.
		if !holder.Honest() {
			err = conn.WaitConcluded(ctx)
			if err != nil {
				errs <- fmt.Errorf("waitinf for dispute: %w", err)
				return
			}
		}

		// Close connection.
		err = conn.Close(ctx)
		if err != nil {
			errs <- fmt.Errorf("closing connection: %w", err)
			return
		}

		holder.Shutdown()
		wg.Done()
	}()

	// Run credential issuer.
	go func() {
		// Connect.
		conn, err := func() (*connection.Connection, error) {
			req, err := issuer.NextConnectionRequest(ctx)
			if err != nil {
				return nil, fmt.Errorf("awaiting next connection request: %w", err)
			}

			// Only accept with correct peer.
			if !req.Peer().Equals(holder.PerunAddress()) {
				return nil, fmt.Errorf("wrong peer: expected %v, got %v", holder, req.Peer())
			}

			conn, err := req.Accept(ctx)
			if err != nil {
				return nil, fmt.Errorf("accepting connection request: %w", err)
			}

			return conn, nil
		}()
		if err != nil {
			errs <- err
			return
		}

		// Issue credential.
		err = func() error {
			req, err := conn.NextCredentialRequest(ctx)
			if err != nil {
				return fmt.Errorf("awaiting next credential request: %w", err)
			}

			// Only accept with correct document and price.
			if err := req.CheckDoc(doc); err != nil {
				return fmt.Errorf("checking document: %w", err)
			} else if req.Price.Cmp(price) < 0 {
				return fmt.Errorf("wrong price: expected %v, got %v", price, req.Price)
			}

			err = req.Accept(ctx)
			if err != nil {
				return fmt.Errorf("accepting credential request: %w", err)
			}

			return nil
		}()
		if err != nil {
			errs <- err
			return
		}

		if conn.Disputed() {
			// If disputed, wait until the channel is concludable.
			err = conn.WaitConcludadable(ctx)
			if err != nil {
				errs <- fmt.Errorf("waiting for channel finalization: %w", err)
				return
			}
		} else {
			// If not disputed, wait until peer finalized the channel.
			err = conn.WaitFinal(ctx)
			if err != nil {
				errs <- fmt.Errorf("waiting for channel finalization: %w", err)
				return
			}
		}

		// Close connection.
		err = conn.Close(ctx)
		if err != nil {
			errs <- fmt.Errorf("closing connection: %w", err)
			return
		}

		issuer.Shutdown()
		wg.Done()
	}()

	// Await result.
	done := make(chan struct{})
	go func() {
		wg.Wait()
		done <- struct{}{}
	}()
	err := func() error {
		select {
		case <-done:
			return nil
		case err := <-errs:
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}()
	require.NoError(err)

	env.LogAccountBalances()
}
