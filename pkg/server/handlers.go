// handlers.go
// This file contains the implementation of the RPC handlers for the key-value store server
// It provides methods to set, get, delete, check existence, and get the length of keys in the store
package server

import (
	"fmt"
)

// Set is an RPC method that sets a key-value pair in the store based on the provided ShardIdx
func (store *KVServer) Set(args *SetArgs, reply *SetReply) error {
	shard := store.shards[args.ShardIdx]
	if shard == nil {
		return fmt.Errorf("shard %d not found", args.ShardIdx)
	}

	shard.mu.Lock()
	defer shard.mu.Unlock()

	shard.data[args.Key] = args.Value

	return nil
}

// Get is an RPC method that retrieves a value by its key from the store based on the provided ShardIdx
// It returns the value and a boolean indicating if the key exists
func (store *KVServer) Get(args *GetArgs, reply *GetReply) error {
	shard := store.shards[args.ShardIdx]
	if shard == nil {
		return fmt.Errorf("shard %d not found", args.ShardIdx)
	}

	shard.mu.RLock()
	defer shard.mu.RUnlock()

	value, exists := shard.data[args.Key]
	if exists {
		reply.Value = value
	}
	reply.Exists = exists

	return nil
}

// Delete is an RPC method that deletes a key from the store based on the provided ShardIdx
// It removes the key from the map if it is there
func (store *KVServer) Delete(args *DeleteArgs, reply *DeleteReply) error {
	shard := store.shards[args.ShardIdx]
	if shard == nil {
		return fmt.Errorf("shard %d not found", args.ShardIdx)
	}

	shard.mu.Lock()
	defer shard.mu.Unlock()

	delete(shard.data, args.Key)

	return nil
}

// Exists is an RPC method that checks if a key exists in the store based on the provided ShardIdx
func (store *KVServer) Exists(args *ExistsArgs, reply *ExistsReply) error {
	shard := store.shards[args.ShardIdx]
	if shard == nil {
		return fmt.Errorf("shard %d not found", args.ShardIdx)
	}

	shard.mu.RLock()
	defer shard.mu.RUnlock()

	_, exists := shard.data[args.Key]
	reply.Exists = exists

	return nil
}

// Length is an RPC method that returns the total number of key-value pairs across all shards
// It sums the lengths of all shards' maps
func (store *KVServer) Length(args *LengthArgs, reply *LengthReply) error {
	reply.Length = 0

	for _, shard := range store.shards {
		shard.mu.RLock()
		reply.Length += len(shard.data)
		shard.mu.RUnlock()
	}

	return nil
}
