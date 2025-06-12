// Package router provides methods for managing a distributed key-value store through a central router
//
// Routers are launched as RPC servers
// Source code and compiled binaries are available in the cmd directory
// The router is designed for high concurrency in both reading and writing operations
// It uses a static routing mechanism based on the hash of the key to determine the appropriate shard
package router
