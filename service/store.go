package service

import (
	"sync"
)

type Store struct {
	mu sync.RWMutex
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Save(key, url string) {
	s.mu.Lock()
	//add to DB
	defer s.mu.Unlock()
}

func (s *Store) Load(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	//url, ok := READ from DB
	//return url, ok
	return "nil", true
}
