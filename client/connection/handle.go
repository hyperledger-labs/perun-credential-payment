package connection

import (
	"context"

	"perun.network/go-perun/channel"
)

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
