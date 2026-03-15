package knownserverstore

import "sync"

type KnownServerStore struct {
	Mutex sync.RWMutex
	Store map[string]struct{}
}

func New() *KnownServerStore {
	return &KnownServerStore{
		Mutex: sync.RWMutex{},
		Store: map[string]struct{}{},
	}
}
