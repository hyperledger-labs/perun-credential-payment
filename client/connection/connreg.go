package connection

import (
	"sync"

	"perun.network/go-perun/channel"
)

type Registry struct {
	mu sync.RWMutex
	r  map[channel.ID]*Connection
}

func NewRegistry() *Registry {
	return &Registry{
		mu: sync.RWMutex{},
		r:  make(map[channel.ID]*Connection),
	}
}

func (r *Registry) Add(conn *Connection) {
	r.mu.Lock()
	r.r[conn.ID()] = conn
	r.mu.Unlock()
}

func (r *Registry) ForID(id channel.ID) (*Connection, bool) {
	r.mu.RLock()
	c, ok := r.r[id]
	r.mu.RUnlock()
	return c, ok
}
