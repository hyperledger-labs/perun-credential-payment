package client

import (
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
	conn, ok := h.connections.ForID(update.State.ID)
	if !ok {
		h.Logf("Update on unknown channel: %x", update.State.ID)
	}

	conn.HandleUpdate(cur, update, responder)
}
