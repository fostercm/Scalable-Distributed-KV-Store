package kvstore

import (
	"sync"
)

type KVStore struct {
	data map[string]string
	mu   sync.Mutex
}

func NewKVStore() *KVStore {
	// Initialize a new Store instance
	return &KVStore{
		data: make(map[string]string),
	}
}
