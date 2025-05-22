package server

import (
	"kvstore/pkg/kvstore"
	"log"
	"net"
	"net/rpc"
)

func Launch(port string) {
	// Register the KVStore service with the RPC server
	kvserver := kvstore.NewKVStore()
	rpcserver := rpc.NewServer()
	rpcserver.Register(kvserver)

	// Start listening for incoming connections on the specified port
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error starting server on port %s: %v", port, err)
	}
	defer listener.Close()

	// Print a message indicating that the server is running
	log.Println("Server is running on port", port)

	// Accept incoming connections in a loop
	for {
		// Accept a new connection
		connection, err := listener.Accept()
		if err == nil {
			// Serve the connection using the RPC server
			go rpcserver.ServeConn(connection)
		} else {
			log.Println("Error accepting connection: ", err)
		}
	}

}
