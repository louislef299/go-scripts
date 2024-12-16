package main

import (
	"log"
	"net"
)

func main() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer listener.Close()

	log.Println("Server is listening on port 8080")
	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error:", err)
			continue
		}

		// Handle client connection in a goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Create a buffer to read data into
	buffer := make([]byte, 1024)
	for {
		// Read data from the client
		n, err := conn.Read(buffer)
		if err != nil && err != net.ErrClosed {
			log.Fatal("Error:", err)
		}

		// Process and use the data (here, we'll just print it)
		log.Printf("Received: %s\n", buffer[:n])
	}
}
