// client.go
// This file contains the required internal client structure and methods
// It provides the basic functionality to connect to a router and initialize a client
package client

import (
	"fmt"
	"kvstore/pkg/router"
	"net/rpc"
)

// Client wraps an RPC client for communication with the router
type Client struct {
	*rpc.Client
	Socket string
}

// NewClient creates a new Client instance connected to the specified address
// It returns a pointer to the Client and an error if connection fails
func NewClient(socket string) (*Client, error) {
	client, err := rpc.Dial("tcp", socket)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to router at %s: %v", socket, err)
	}

	newClient := &Client{
		Client: client,
		Socket: socket,
	}

	return newClient, nil
}

// getShardClient retrieves the shard client for a given key
// It queries the router and establishes an RPC connection to the appropriate shard server
// It returns the shard client, the shard index, and an error if any occur
func (c *Client) getShardClient(key string) (*Client, int, error) {
	args := &router.GetRouteArgs{Key: key}
	reply := &router.GetRouteReply{}
	err := c.Call("StaticShardRouter.GetRoute", args, reply)
	if err != nil {
		return nil, 0, fmt.Errorf("route error for key %s: %v", key, err)
	}

	shardClient, err := NewClient(reply.Socket)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create shard client for socket %s: %v", reply.Socket, err)
	}

	return shardClient, reply.ShardIdx, nil
}

// getAllSockets retrieves all sockets managed by the router
// It returns a slice of strings containing the socket addresses and an error if any occur
func (c *Client) getAllSockets() ([]string, error) {
	args := &router.GetAllSocketsArgs{}
	reply := &router.GetAllSocketsReply{}
	err := c.Call("StaticShardRouter.GetAllSockets", args, reply)
	if err != nil {
		return nil, fmt.Errorf("unable to get all sockets: %v", err)
	}

	return reply.Sockets, nil
}
