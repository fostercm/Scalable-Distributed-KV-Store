// rpc_types.go
// This file contains the RPC types used for communication between the router and clients
package router

// GetRouteArgs and GetRouteReply are used for the GetRoute RPC method
// This method retrieves the route for a given key
// This RPC is used for all routing operations
type GetRouteArgs struct {
	Key string
}

type GetRouteReply struct {
	Socket   string
	ShardIdx int
}

// GetAllSocketsArgs and GetAllSocketsReply are used for the GetAllSockets RPC method
// This method retrieves all registered shard sockets
// This RPC is currently used for getting the overall size of the key-value store
type GetAllSocketsArgs struct{}

type GetAllSocketsReply struct {
	Sockets []string
}

// RegisterServerArgs and RegisterServerReply are used for the RegisterServer RPC method
// This method allows a new server to register itself with the router
// It takes the address, port, and number of shards on the server as arguments
type RegisterServerArgs struct {
	Address   string
	Port      int
	NumShards int
}

type RegisterServerReply struct{}
