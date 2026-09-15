package network

import (
	"OnePiece/core"
	"fmt"
	"log"
	"net"
)

func ListenTcp(app *core.App) {

	listener, err := net.Listen("tcp", "127.0.0.1:3003")
	if err != nil {
		log.Fatalf("Failed to bind to port: %v", err)
	}
	defer listener.Close()
	fmt.Println("Server running on http://127.0.0.1:3003...")

	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept network: %v", err)
			continue
		}

		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("Failed to read from connection: %v", err)
			conn.Close()
			continue
		}

		data := buf[:n]
		fmt.Printf("Received %d bytes: %s\n", n, string(data))
	}

}
