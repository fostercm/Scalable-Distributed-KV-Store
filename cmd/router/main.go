package main

import (
	"flag"
	"kvstore/pkg/router"
	"log"
	"net"
	"net/rpc"
)

func main() {
	// Get command-line arguments
	port := flag.String("port", "8080", "Port to run the server on")
	flag.Parse()

	// Register the router with the RPC server
	routeController := router.NewRouter()
	rpcserver := rpc.NewServer()
	rpcserver.Register(routeController)

	// Start listening for incoming connections on the specified port
	listener, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("Error starting router on port %s: %v", *port, err)
	}
	defer listener.Close()

	// Print a message indicating that the server is running
	log.Println("Router is running on port", *port)

	// Accept and serve incoming connections
	for {
		connection, err := listener.Accept()
		if err == nil {
			go rpcserver.ServeConn(connection)
		} else {
			log.Println("Error accepting connection: ", err)
		}
	}
}
