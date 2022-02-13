package channel

import (
	"sync"

	"perun.network/go-perun/channel"
)

type Registry struct {
	mu sync.RWMutex
	r  map[channel.ID]*Channel
}

func NewRegistry() *Registry {
	return &Registry{
		mu: sync.RWMutex{},
		r:  make(map[channel.ID]*Channel),
	}
}

func (r *Registry) Add(ch *Channel) {
	r.mu.Lock()
	r.r[ch.ID()] = ch
	r.mu.Unlock()
}

func (r *Registry) ForID(id channel.ID) (*Channel, bool) {
	r.mu.RLock()
	c, ok := r.r[id]
	r.mu.RUnlock()
	return c, ok
}
