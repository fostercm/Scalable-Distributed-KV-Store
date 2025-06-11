// operations.go
// This file contains the implementation of client-side operations such as Set, Get, Delete, Exists, and Length
// It uses the server package for RPC calls to the appropriate shard based on the key's routing
package client

import (
	"kvstore/pkg/server"
	"log"
)

// Set routes a key to the appropriate shard and sets its value
// It returns an error if routing or set RPC call fails
func (c *Client) Set(key string, value string) error {
	shardClient, shardIdx, err := c.getShardClient(key)
	if err != nil {
		return err
	}

	args := &server.SetArgs{Key: key, Value: value, ShardIdx: shardIdx}
	reply := &server.SetReply{}

	err = shardClient.Call("KVServer.Set", args, reply)

	return err
}

// Get retrieves the value for a given key from the appropriate shard
// It returns the value, a boolean indicating if the key exists, and an error if any occur
func (c *Client) Get(key string) (string, bool, error) {
	shardClient, shardIdx, err := c.getShardClient(key)
	if err != nil {
		return "", false, err
	}

	args := &server.GetArgs{Key: key, ShardIdx: shardIdx}
	reply := &server.GetReply{}

	err = shardClient.Call("KVServer.Get", args, reply)
	if err != nil {
		return "", false, err
	}

	return reply.Value, reply.Exists, nil
}

// Delete removes a key from the appropriate shard
// It returns an error if the delete RPC call fails
func (c *Client) Delete(key string) error {
	shardClient, shardIdx, err := c.getShardClient(key)
	if err != nil {
		return err
	}

	args := &server.DeleteArgs{Key: key, ShardIdx: shardIdx}
	reply := &server.DeleteReply{}

	err = shardClient.Call("KVServer.Delete", args, reply)

	return err
}

// Exists checks if a key exists in the appropriate shard
// It returns a boolean indicating if the key exists and an error if any occur
func (c *Client) Exists(key string) (bool, error) {
	shardClient, shardIdx, err := c.getShardClient(key)
	if err != nil {
		return false, err
	}

	args := &server.ExistsArgs{Key: key, ShardIdx: shardIdx}
	reply := &server.ExistsReply{}

	err = shardClient.Call("KVServer.Exists", args, reply)
	if err != nil {
		return false, err
	}

	return reply.Exists, nil
}

// Length calculates the total number of keys across all shards
// Shards that are unreachable or return an error are logged but do not affect the total count
func (c *Client) Length() (int, error) {
	length := 0

	sockets, err := c.getAllSockets()
	if err != nil {
		log.Println("Error getting all sockets:", err)
		return 0, err
	}

	for _, socket := range sockets {
		serverClient, err := NewClient(socket)
		if err != nil {
			log.Println("Error creating client for socket:", socket, err)
			continue
		}

		args := &server.LengthArgs{}
		reply := &server.LengthReply{}

		err = serverClient.Call("KVServer.Length", args, reply)
		if err != nil {
			log.Println("Error calling Length on socket:", socket, err)
			continue
		}

		length += reply.Length
		serverClient.Close()
	}

	return length, nil
}
