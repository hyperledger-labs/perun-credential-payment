package client

import (
	"bytes"
	"context"
	"fmt"

	"github.com/perun-network/verifiable-credential-payment/app/data"
	"github.com/perun-network/verifiable-credential-payment/client/connection"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/client"
)

type handler struct {
	*Client
}

func (h *handler) HandleProposal(p client.ChannelProposal, r *client.ProposalResponder) {
	lp, ok := p.(*client.LedgerChannelProposal)
	if !ok {
		h.Logf("invalid proposal type: %T", p)
		return
	}
	h.channelProposals <- connection.NewChannelProposal(lp, r)
}

func (h *handler) HandleUpdate(cur *channel.State, update client.ChannelUpdate, responder *client.UpdateResponder) {
	switch nextData := update.State.Data.(type) {
	case *data.Offer:
		// Check if we are requested issuer.
		if acc := h.Address(); !bytes.Equal(acc[:], nextData.Issuer[:]) {
			h.Logf("Received offer for different issuer: %v", nextData.Issuer)
			return
		}

		// Ask user for response.
		conn := h.connections.ForID(update.State.ID)
		response := conn.AddCredentialRequest(nextData.DataHash, nextData.Price)

		// Send response.
		r := <-response
		switch r.(type) {
		case *connection.CredentialRequestResponseAccept:
			err := responder.Accept(r.Context())
			if err != nil {
				r.Result() <- fmt.Errorf("accepting update: %w", err)
				return
			}

			// Proceed asynchronously because handler must return before new updates can be processed.
			go func() {
				err = conn.IssueCredential(r.Context(), nextData, h.perunClient.Account)
				if err != nil {
					r.Result() <- fmt.Errorf("issueing credential: %w", err)
					return
				}

				r.Result() <- nil
			}()

		default:
			panic(fmt.Sprintf("unsupported type: %T", r))
		}

	case *data.Cert:
		prevData := cur.Data.(*data.Offer)
		conn := h.connections.ForID(update.State.ID)

		// The app logic ensures that the signature is valid.
		conn.AddSignature(nextData.Signature[:], prevData.DataHash, prevData.Issuer)

		if h.honest {
			// We are honest. We received the signature. We accept the update.
			err := responder.Accept(context.TODO())
			if err != nil {
				h.Logf("Error accepting update: %v", err)
				return
			}
		} else {
			// We are dishonest. We received the signature, but we reject the update.
			err := responder.Reject(context.TODO(), "Won't pay!")
			if err != nil {
				h.Logf("Error rejecting update: %v", err)
				return
			}
		}

	case *data.DefaultData:
		// Always accept update. The app logic ensures that the balances do not
		// change.
		err := responder.Accept(context.TODO())
		if err != nil {
			h.Logf("Error accepting update: %v", err)
			return
		}

	default:
		h.Logf("Unexpected data type: %T", nextData)

	}
}
