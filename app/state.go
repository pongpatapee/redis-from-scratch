package main

import (
	"sync"
)

type RedisStringDB struct {
	store map[string]string
	mu    sync.RWMutex
}

func NewRedisStringDB() *RedisStringDB {
	return &RedisStringDB{store: make(map[string]string)}
}

func (db *RedisStringDB) Set(key string, val string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.store[key] = val
}

func (db *RedisStringDB) Del(key string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.store, key)
}

func (db *RedisStringDB) Get(key string) (string, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()
	value, exist := db.store[key]

	return value, exist
}

var StringDB = NewRedisStringDB()
