package connection

import (
	"context"
	"fmt"

	"github.com/perun-network/perun-credential-payment/app/data"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/client"
)

func (conn *Connection) HandleUpdate(cur *channel.State, update client.ChannelUpdate, responder *client.UpdateResponder) {
	switch nextData := update.State.Data.(type) {
	case *data.Offer:
		conn.handleOffer(nextData, responder)

	case *data.Cert:
		curData := cur.Data.(*data.Offer)
		conn.handleCert(curData, nextData, responder)

	case *data.DefaultData:
		// Always accept update. The app logic ensures that the balances do not
		// change.
		err := responder.Accept(context.TODO())
		if err != nil {
			conn.Log().Warnf("Error accepting update: %v", err)
			return
		}

	default:
		conn.Log().Warnf("Unexpected data type: %T", nextData)

	}
}

func (conn *Connection) handleOffer(offer *data.Offer, responder *client.UpdateResponder) {
	// Forward the request and get response.
	response := conn.addCredentialRequest(offer)
	r := <-response

	// Send response.
	switch r.(type) {
	case *CredentialRequestResponseAccept:
		err := responder.Accept(r.Context())
		if err != nil {
			r.Result() <- fmt.Errorf("accepting update: %w", err)
			return
		}

		r.Result() <- nil

	default:
		panic(fmt.Sprintf("unsupported type: %T", r))
	}
}

func (conn *Connection) handleCert(curData *data.Offer, nextData *data.Cert, responder *client.UpdateResponder) {
	// The app logic ensures that the signature is valid.
	conn.addSignature(nextData.Signature[:], curData.DataHash, curData.Issuer, responder)
}

type EventHandler struct {
	*Connection
}

func NewEventHandler(conn *Connection) *EventHandler {
	return &EventHandler{Connection: conn}
}

func (h *EventHandler) HandleAdjudicatorEvent(e channel.AdjudicatorEvent) {
	switch e := e.(type) {
	case *channel.RegisteredEvent:
		h.disputed.SetValue(true)
	case *channel.ProgressedEvent:
		go func() {
			err := e.TimeoutV.Wait(context.TODO())
			if err != nil {
				h.Log().Warnf("waiting for timeout: %v", err)
				return
			}
			h.concludable.SetValue(true)
		}()
	case *channel.ConcludedEvent:
		h.concluded.SetValue(true)
	}
}
