// rpc_types.go
// This file contains the RPC types used for the key-value store server
// The shard index provided by the router is used to determine which shard to access
package server

// The Set RPC method is used to set a key-value pair in the store
type SetArgs struct {
	Key      string
	Value    string
	ShardIdx int
}

type SetReply struct{}

// The Get RPC method is used to retrieve a value by its key
type GetArgs struct {
	Key      string
	ShardIdx int
}

type GetReply struct {
	Value  string
	Exists bool
}

// The Delete RPC method is used to delete a key from the store
type DeleteArgs struct {
	Key      string
	ShardIdx int
}

type DeleteReply struct{}

// The Exists RPC method checks if a key exists in the store
type ExistsArgs struct {
	Key      string
	ShardIdx int
}

type ExistsReply struct {
	Exists bool
}

// The Length RPC method returns the number of keys in the store
type LengthArgs struct{}

type LengthReply struct {
	Length int
}
