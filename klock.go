package klock

import "sync"

type Klock struct {
	locks map[interface{}]*sync.Mutex
	mu    sync.Mutex
}

func New() *Klock {
	return &Klock{
		locks: map[interface{}]*sync.Mutex{},
	}
}

func (k *Klock) Lock(key interface{}) {
	k.mu.Lock()
	lc := k.locks[key]
	if lc == nil {
		lc = &sync.Mutex{}
		k.locks[key] = lc
	}
	k.mu.Unlock()
	lc.Lock()
}

func (k *Klock) Unlock(key interface{}) {
	k.mu.Lock()
	defer k.mu.Unlock()
	lc := k.locks[key]
	if lc == nil {
		panic("unlock at nil lock")
	}
	lc.Unlock()
}
