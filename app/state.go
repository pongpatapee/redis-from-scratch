package main

import (
	"sync"
)

type Store struct {
	store map[string]string
	mu    sync.Mutex
}

func (s *Store) Set(key string, val string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.store[key] = val
}

func (s *Store) Del(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.store, key)
}

func (s *Store) Get(key string) (string, bool) {
	value, exist := s.store[key]

	return value, exist
}

var StringStore = Store{store: make(map[string]string)}
