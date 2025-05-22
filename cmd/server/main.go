package main

import (
	"flag"
	"kvstore/pkg/server"
)

func main() {
	// Define the command-line flags
	port := flag.String("port", "8080", "Port to run the server on")
	flag.Parse()

	// Launch the server with the specified port
	server.Launch(":" + *port)
}
