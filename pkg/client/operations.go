package client

import (
	"kvstore/pkg/kvstore"
)

func (c *Client) Set(key, value string) error {
	// Create a SetArgs instance with the provided key and value
	args := &kvstore.SetArgs{Key: key, Value: value}
	reply := &kvstore.SetReply{}

	// Call the Set method on the kvclient
	err := c.Call("KVStore.Set", args, reply)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Get(key string) (string, bool, error) {
	// Create a GetArgs instance with the provided key
	args := &kvstore.GetArgs{Key: key}
	reply := &kvstore.GetReply{}

	// Call the Get method on the kvclient
	err := c.Call("KVStore.Get", args, reply)
	if err != nil {
		return "", false, err
	}

	return reply.Value, reply.Exists, nil
}

func (c *Client) Delete(key string) error {
	// Create a DeleteArgs instance with the provided key
	args := &kvstore.DeleteArgs{Key: key}
	reply := &kvstore.DeleteReply{}

	// Call the Delete method on the kvclient
	err := c.Call("KVStore.Delete", args, reply)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Exists(key string) (bool, error) {
	// Create a ExistsArgs instance
	args := &kvstore.ExistsArgs{Key: key}
	reply := &kvstore.ExistsReply{}

	// Call the Exists method on the kvclient
	err := c.Call("KVStore.Exists", args, reply)
	if err != nil {
		return false, err
	}

	return reply.Exists, nil
}

func (c *Client) Length() (int, error) {
	// Create a LengthArgs instance
	args := &kvstore.LengthArgs{}
	reply := &kvstore.LengthReply{}

	// Call the Length method on the kvclient
	err := c.Call("KVStore.Length", args, reply)
	if err != nil {
		return 0, err
	}

	return reply.Length, nil
}
