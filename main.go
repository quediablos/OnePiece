package main

import (
	"OnePiece/core"
	"fmt"
	"log"
	"net"
	"runtime"

	"OnePiece/connection"
)

func main() {

	//Start the app data.
	app := NewApp()
	_ = app

	runtime.GOMAXPROCS(1)

	// 2. Bind to a port and listen for TCP traffic
	listener, err := net.Listen("tcp", "127.0.0.1:3003")
	if err != nil {
		log.Fatalf("Failed to bind to port: %v", err)
	}
	defer listener.Close()
	fmt.Println("Single-threaded server running on http://127.0.0.1:3003...")

	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		req, err := connection.ReadHTTPRequest(conn)
		if err != nil {
			log.Printf("Failed to parse request: %v", err)
			conn.Close()
			continue
		}

		fmt.Printf("Processing a request — Method: %s, Path: %s, Query: %s\n", req.Method, req.Path, req.Query)
		fmt.Printf("Body: %s\n", string(req.Body))

		operation, resourceId, err := req.ParseURL()

		if operation == core.Lock {

			lock := core.CheckForLock(resourceId, app.Locks)

			if lock != nil {

				fmt.Printf("Lock found for resourceId: %s\n", resourceId)
			}
		}

		response := "HTTP/1.1 200 OK\r\n" +
			"Content-Type: text/plain\r\n" +
			"Content-Length: 13\r\n" +
			"\r\n" +
			"Hello, World!"

		conn.Write([]byte(response))
		conn.Close()
		fmt.Println("Finished processing request.")
	}
}
