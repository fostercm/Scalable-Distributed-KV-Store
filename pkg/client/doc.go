// Package client provides a client-side abstraction for interacting with a distributed key-value store
//
// It exposes two main functionalities:
//  1. Routing: Abstracts the connection to the central router
//  2. Operations: Provides methods to perform CRUD operations on the key-value store including:
//     - Set
//     - Get
//     - Delete
//     - Exists
//     - Length
//
// # Clients are created using NewClient(address) which connects to the specified router address
//
// Example usage:
//
//	client, err := client.NewClient("localhost:1234")
//	client.Set("key1", "value1")
//	value, exists, err := client.Get("key1")
//
// The client is designed to be hardly differentiated from a standard Go map
package client
