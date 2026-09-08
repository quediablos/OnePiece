package main

import (
	"fmt"
	"net"
	"time"
)

// Global map to hold active connections by ID
var clients = make(map[string]net.Conn)

func handleClient(conn net.Conn) {
	clientID := "client-1" // Use a unique ID or address in real apps
	clients[clientID] = conn

	fmt.Println("Connected to client:", conn.RemoteAddr())

	// Wait in a goroutine or let the main/other logic handle the delay
	go delayedResponse(clientID)
}

func delayedResponse(clientID string) {
	// Hold the connection for 5 seconds
	time.Sleep(5 * time.Second)

	conn, exists := clients[clientID]
	if !exists {
		fmt.Println("Client not found")
		return
	}

	// Send the response later
	_, err := conn.Write([]byte("Hello from the server after a delay!\n"))
	if err != nil {
		fmt.Println("Error writing:", err)
	}

	// Close the connection after responding
	conn.Close()
	delete(clients, clientID)
}

func main_2_rename() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err)
			continue
		}

		go handleClient(conn)
	}
}
