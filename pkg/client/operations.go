// operations.go
// This file contains the implementation of client-side operations such as Set, Get, Delete, Exists, and Length
// It uses the server package for RPC calls to the appropriate shard based on the key's routing
package client

import (
	"fmt"
	"kvstore/pkg/server"
)

// Set routes a key to the appropriate shard and sets its value
// It returns an error if routing or set RPC call fails
func (c *Client) Set(key string, value string) error {
	shardClient, shardIdx, err := c.getShardClient(key)
	if err != nil {
		return fmt.Errorf("failed to get shard client for key %s: %v", key, err)
	}

	args := &server.SetArgs{Key: key, Value: value, ShardIdx: shardIdx}
	reply := &server.SetReply{}

	err = shardClient.Call("KVServer.Set", args, reply)
	if err != nil {
		return fmt.Errorf("failed to set value for key %s at socket %s and shard index %d: %v", key, shardClient.Socket, shardIdx, err)
	}

	return nil
}

// Get retrieves the value for a given key from the appropriate shard
// It returns the value, a boolean indicating if the key exists, and an error if any occur
func (c *Client) Get(key string) (string, bool, error) {
	shardClient, shardIdx, err := c.getShardClient(key)
	if err != nil {
		return "", false, fmt.Errorf("failed to get shard client for key %s: %v", key, err)
	}

	args := &server.GetArgs{Key: key, ShardIdx: shardIdx}
	reply := &server.GetReply{}

	err = shardClient.Call("KVServer.Get", args, reply)
	if err != nil {
		return "", false, fmt.Errorf("failed to get value for key %s at socket %s and shard index %d: %v", key, shardClient.Socket, shardIdx, err)
	}

	return reply.Value, reply.Exists, nil
}

// Delete removes a key from the appropriate shard
// It returns an error if the delete RPC call fails
func (c *Client) Delete(key string) error {
	shardClient, shardIdx, err := c.getShardClient(key)
	if err != nil {
		return fmt.Errorf("failed to get shard client for key %s: %v", key, err)
	}

	args := &server.DeleteArgs{Key: key, ShardIdx: shardIdx}
	reply := &server.DeleteReply{}

	err = shardClient.Call("KVServer.Delete", args, reply)
	if err != nil {
		return fmt.Errorf("failed to delete key %s at socket %s and shard index %d: %v", key, shardClient.Socket, shardIdx, err)
	}

	return nil
}

// Exists checks if a key exists in the appropriate shard
// It returns a boolean indicating if the key exists and an error if any occur
func (c *Client) Exists(key string) (bool, error) {
	shardClient, shardIdx, err := c.getShardClient(key)
	if err != nil {
		return false, fmt.Errorf("failed to get shard client for key %s: %v", key, err)
	}

	args := &server.ExistsArgs{Key: key, ShardIdx: shardIdx}
	reply := &server.ExistsReply{}

	err = shardClient.Call("KVServer.Exists", args, reply)
	if err != nil {
		return false, fmt.Errorf("failed to check existence of key %s at socket %s and shard index %d: %v", key, shardClient.Socket, shardIdx, err)
	}

	return reply.Exists, nil
}

// Length calculates the total number of keys across all shards
// Shards that are unreachable or return an error are logged but do not affect the total count
// Errors encountered during the Length operation are aggregated and returned without stopping the operation
func (c *Client) Length() (int, error) {
	length := 0

	sockets, err := c.getAllSockets()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve sockets: %v", err)
	}

	overallErr := fmt.Errorf("errors encountered during Length operation: ")
	errFlag := false

	for _, socket := range sockets {
		serverClient, err := NewClient(socket)
		if err != nil {
			overallErr = fmt.Errorf("%w\nSocket=%s, SubError=%v", overallErr, socket, err)
			errFlag = true
			continue
		}

		args := &server.LengthArgs{}
		reply := &server.LengthReply{}

		err = serverClient.Call("KVServer.Length", args, reply)
		if err != nil {
			overallErr = fmt.Errorf("%w\nSocket=%s, SubError=%v", overallErr, socket, err)
			errFlag = true
			continue
		}

		length += reply.Length
		serverClient.Close()
	}

	if errFlag {
		return length, overallErr
	}
	return length, nil
}
