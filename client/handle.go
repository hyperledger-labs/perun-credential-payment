package client

import (
	clientchannel "github.com/perun-network/perun-credential-payment/client/channel"
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
	h.channelProposals <- clientchannel.NewChannelProposal(lp, r)
}

func (h *handler) HandleUpdate(cur *channel.State, update client.ChannelUpdate, responder *client.UpdateResponder) {
	ch, ok := h.channels.ForID(update.State.ID)
	if !ok {
		h.Logf("Update on unknown channel: %x", update.State.ID)
	}

	ch.HandleUpdate(cur, update, responder)
}
