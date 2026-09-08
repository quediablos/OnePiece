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
				//Another client already locked the resource, hold this client until the lock is released.
				fmt.Printf("Lock found for resourceId: %s\n", resourceId)
				core.RequestLock(conn, resourceId, app.LockRequests)
			} else {

				//No lock is acquired for the resource id yet, acquire the lock, and finish the connectin.
				fmt.Printf("Lock not found for resourceId: %s\n", resourceId)
				core.AcquireLock(resourceId, app.Locks)

				response := "HTTP/1.1 200 OK\r\n" +
					"Content-Type: text/plain\r\n" +
					"Content-Length: 50\r\n" +
					"\r\n" +
					"Acquired lock for resourceId: " + resourceId
				connection.WriteResponseHttp(conn, response)
			}
		} else if operation == core.Unlock {

			waitingOne := core.ReleaseLock(resourceId, app.Locks, app.LockRequests)

			if waitingOne != nil {
				responseWaitingOne := "HTTP/1.1 200 OK\r\n" +
					"Content-Type: text/plain\r\n" +
					"Content-Length: 50\r\n" +
					"\r\n" +
					"Acquired lock for resourceId: " + resourceId
				connection.WriteResponseHttp(waitingOne, responseWaitingOne)
			}

			response := "HTTP/1.1 200 OK\r\n" +
				"Content-Type: text/plain\r\n" +
				"Content-Length: 50\r\n" +
				"\r\n" +
				"Released lock for resourceId: " + resourceId
			connection.WriteResponseHttp(conn, response)
		}
	}
}
