package main_test

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"testing"

	"github.com/perun-network/perun-credential-payment/app"
	"github.com/perun-network/perun-credential-payment/client"
	"github.com/perun-network/perun-credential-payment/client/connection"
	"github.com/perun-network/perun-credential-payment/test"
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
	env := test.Setup(t)
	env.LogAccountBalances()
	wg, errs := sync.WaitGroup{}, make(chan error)
	wg.Add(2)
	holder, issuer := env.Holder, env.Issuer

	doc := []byte("Perun/Bosch: SSI Credential Payment")
	balance := test.EthToWei(big.NewFloat(5))
	price := test.EthToWei(big.NewFloat(1))

	// Run credential holder.
	go func() {
		err := runCredentialHolder(
			ctx,
			holder,
			issuer,
			balance,
			doc,
			price,
			honestHolder,
		)
		if err != nil {
			errs <- fmt.Errorf("running credential holder: %w", err)
			return
		}

		wg.Done()
	}()

	// Run credential issuer.
	go func() {
		err := runCredentialIssuer(
			ctx,
			issuer,
			holder,
			doc,
			price,
		)
		if err != nil {
			errs <- fmt.Errorf("running credential issuer: %w", err)
			return
		}

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

func runCredentialHolder(
	ctx context.Context,
	holder *client.Client,
	issuer *client.Client,
	balance *big.Int,
	doc []byte,
	price *big.Int,
	honest bool,
) error {
	// Open channel.
	ch, err := holder.OpenChannel(ctx, issuer.PerunAddress(), balance)
	if err != nil {
		return fmt.Errorf("proposing channel: %w", err)
	}

	// Buy credential.
	{
		// Request credential.
		asyncCred, err := ch.RequestCredential(ctx, doc, price, issuer.EthAddress())
		if err != nil {
			return fmt.Errorf("requesting credential: %w", err)
		}

		// Wait for the transaction issueing the credential.
		resp, err := asyncCred.Await(ctx)
		if err != nil {
			return fmt.Errorf("awaiting credential: %w", err)
		}

		cred := app.Credential{
			Document:  doc,
			Signature: resp.Signature,
		}
		holder.Logf("Obtained credential: %v", cred.String())

		// The issuer is waiting for us to complete the transaction.
		// If we are honest, we accept. If we are dishonest, we reject.
		if honest {
			err := resp.Accept(ctx)
			if err != nil {
				return fmt.Errorf("accepting transaction: %w", err)
			}
		} else {
			err := resp.Reject(ctx, "Won't pay!")
			if err != nil {
				return fmt.Errorf("rejecting transaction: %w", err)
			}

			// We wait for the dispute to be resolved.
			err = ch.WaitConcludadable(ctx)
			if err != nil {
				return fmt.Errorf("waiting for dispute resolution: %w", err)
			}
		}
	}

	// Close connection.
	err = ch.Close(ctx)
	if err != nil {
		return fmt.Errorf("closing connection: %w", err)
	}

	return nil
}

func runCredentialIssuer(
	ctx context.Context,
	issuer *client.Client,
	holder *client.Client,
	doc []byte,
	price *big.Int,
) error {
	// Connect.
	conn, err := func() (*connection.Channel, error) {
		// Read next connection request.
		req, err := issuer.NextConnectionRequest(ctx)
		if err != nil {
			return nil, fmt.Errorf("awaiting next connection request: %w", err)
		}

		// Check peer.
		if !req.Peer().Equals(holder.PerunAddress()) {
			return nil, fmt.Errorf("wrong peer: expected %v, got %v", holder, req.Peer())
		}

		// Accept.
		conn, err := req.Accept(ctx)
		if err != nil {
			return nil, fmt.Errorf("accepting connection request: %w", err)
		}

		return conn, nil
	}()
	if err != nil {
		return fmt.Errorf("connecting: %w", err)
	}

	// Issue credential.
	err = func() error {
		// Read next credential request.
		req, err := conn.NextCredentialRequest(ctx)
		if err != nil {
			return fmt.Errorf("awaiting next credential request: %w", err)
		}

		// Check document and price.
		if err := req.CheckDoc(doc); err != nil {
			return fmt.Errorf("checking document: %w", err)
		} else if err := req.CheckPrice(price); err != nil {
			return fmt.Errorf("checking price: %w", err)
		}

		// Issue credential.
		err = req.IssueCredential(ctx, issuer.Account())
		if err != nil {
			return fmt.Errorf("issueing credential: %w", err)
		}

		return nil
	}()
	if err != nil {
		return fmt.Errorf("issueing credential: %w", err)
	}

	// Wait until channel is concludable.
	err = conn.WaitConcludadable(ctx)
	if err != nil {
		return fmt.Errorf("waiting for channel finalization: %w", err)
	}

	// Close connection.
	err = conn.Close(ctx)
	if err != nil {
		return fmt.Errorf("closing connection: %w", err)
	}

	return nil
}
