// kvstore.go
// This file contains the core of a key-value server with sharding
// It provides a struct for the KVServer and a method to initialize it
package server

import (
	"sync"
)

// A Shard in the key-value store is a thread-safe map
type Shard struct {
	data map[string]string
	mu   sync.RWMutex
}

// The KVServer is a list of shards
type KVServer struct {
	shards []*Shard
}

// NewShard initializes an empty Shard instance
func NewShard() *Shard {
	return &Shard{
		data: make(map[string]string),
	}
}

// NewKVServer initializes a new KVServer with the specified number of shards
func NewKVServer(numShards int) *KVServer {
	shards := make([]*Shard, numShards)
	for i := range numShards {
		shards[i] = NewShard()
	}

	return &KVServer{
		shards: shards,
	}
}
