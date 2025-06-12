// router.go
// This file contains the implementation of a central shard router
// It provides structs and methods to route requests to the appropriate shard based on a key
// It uses a static routing mechanism based on the hash of the key
package router

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"sync"

	"github.com/cespare/xxhash/v2"
)

// The ShardRoute struct contains the necessary information to route a request to a specific shard
// The address and port are necessary for RPC communication with a shard server
// The shard index is used to identify which shard on a server the request should be routed to
type ShardRoute struct {
	Socket   string
	ShardIdx int
}

// The StaticShardRouter struct contains the routing information for all shards
// It holds a slice of routes to each shard and the total number of shards
type StaticShardRouter struct {
	Routes []*ShardRoute
	mu     sync.RWMutex
}

// NewRouter initializes a new StaticShardRouter with an empty route list and zero shards
func NewRouter() *StaticShardRouter {
	return &StaticShardRouter{
		Routes: make([]*ShardRoute, 0),
	}
}

// GetRoute is an RPC method that retrieves the route for a given key
// It calculates the 64-bit hash of the key and determines the appropriate shard based on the hash
// Thread-safe access is ensured using a read mutex
// The reply contains the socket and shard index for the requested key
func (r *StaticShardRouter) GetRoute(args *GetRouteArgs, reply *GetRouteReply) error {
	// xxhash is used for fast, non-cryptographic hashing of keys for simple and efficient routing
	hash := xxhash.Sum64String(args.Key)
	routeIdx := int(hash % uint64(len(r.Routes)))

	r.mu.RLock()
	route := r.Routes[routeIdx]
	r.mu.RUnlock()
	if route == nil {
		return fmt.Errorf("no route found for key %s", args.Key)
	}

	reply.Socket = route.Socket
	reply.ShardIdx = route.ShardIdx

	return nil
}

// GetAllSockets is an RPC method that retrieves all registered shard sockets
// It returns a slice of strings, each representing the address and port of a shard
// Thread-safe access is ensured using a read mutex
func (r *StaticShardRouter) GetAllSockets(args *GetAllSocketsArgs, reply *GetAllSocketsReply) error {
	reply.Sockets = make([]string, 0)

	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, route := range r.Routes {
		if !slices.Contains(reply.Sockets, route.Socket) {
			reply.Sockets = append(reply.Sockets, route.Socket)
		}
	}

	return nil
}

// RegisterServer is an RPC method that allows a new server to register itself with the router
// It takes the address, port, and number of shards on the server as arguments
// Thread-safe access is ensured using a write mutex
func (r *StaticShardRouter) RegisterServer(args *RegisterServerArgs, reply *RegisterServerReply) error {
	if args.Port < 0 || args.Port > 65535 {
		return fmt.Errorf("valid port numbers are 0-65535, got: %d", args.Port)
	}
	if args.NumShards <= 0 {
		return fmt.Errorf("number of shards must be greater than 0, got: %d", args.NumShards)
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range args.NumShards {
		route := &ShardRoute{
			Socket:   args.Address + ":" + strconv.Itoa(args.Port),
			ShardIdx: i,
		}

		r.Routes = append(r.Routes, route)
	}

	log.Println(
		"Registered new server:",
		"\n\tAddress: ", args.Address,
		"\n\tPort: ", args.Port,
		"\n\tShards: ", args.NumShards,
		"\n\tTotal Shards: ", len(r.Routes),
	)

	return nil
}
