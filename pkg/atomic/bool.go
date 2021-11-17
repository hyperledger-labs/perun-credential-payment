package atomic

import "sync"

type Bool struct {
	mux   sync.RWMutex
	value bool
}

func NewBool(value bool) *Bool {
	return &Bool{
		mux:   sync.RWMutex{},
		value: value,
	}
}

func (a *Bool) SetValue(v bool) {
	a.mux.Lock()
	a.value = v
	a.mux.Unlock()
}

func (a *Bool) Value() bool {
	a.mux.RLock()
	defer a.mux.RUnlock()
	return a.value
}
