package client

import (
	"net/rpc"
)

type Client struct {
	*rpc.Client
}

func NewClient(address string) (*Client, error) {
	// Connect to the server at the specified address
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	// Return a new Client instance with the connected kvclient
	return &Client{Client: client}, nil
}
