package main

import (
	"flag"
	"kvstore/pkg/router"
	"kvstore/pkg/server"
	"log"
	"net"
	"net/rpc"
	"strconv"
)

func main() {
	// Define the command-line flags
	address := flag.String("address", "localhost", "Address to bind the server to")
	port := flag.String("port", "8081", "Port to run the server on")
	numShards := flag.Int("numShards", 4, "Number of shards to use")
	routerSocket := flag.String("routerSocket", "", "Socket address of the router")
	flag.Parse()

	// Register the KVStore service with the RPC server
	kvserver := server.NewKVServer(*numShards)
	rpcserver := rpc.NewServer()
	rpcserver.Register(kvserver)

	// Connect with the router if a socket is provided
	if *routerSocket == "" {
		log.Println("Please provide a router socket address using the -routerSocket flag")
		return
	}
	conn, err := rpc.Dial("tcp", *routerSocket)
	if err != nil {
		log.Println("Error connecting to router:", err)
		return
	}

	numPort, err := strconv.Atoi(*port)
	if err != nil {
		log.Println("Invalid port number:", *port)
		return
	}
	conn.Call("StaticShardRouter.RegisterServer", &router.RegisterServerArgs{
		Address:   *address,
		Port:      numPort,
		NumShards: *numShards,
	}, nil)
	conn.Close()

	// Start listening for incoming connections on the specified port
	listener, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Println("Error starting server: ", err)
	}
	defer listener.Close()

	// Print a message indicating that the server is running
	log.Println("Server is running on port", *port)
	log.Println("Number of shards:", *numShards)

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
