package network

import (
	"OnePiece/core"
	"OnePiece/message"
	"fmt"
	"log"
	"net"
)

func ListenHttp(app *core.App) {

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

		req, err := ReadHTTPRequest(conn)
		if err != nil {
			log.Printf("Failed to parse request: %v", err)
			conn.Close()
			continue
		}

		operation, resourceId, extraParams, err := req.ParseURL()
		operationData := core.OperationData{
			Operation:   operation,
			ResourceId:  resourceId,
			ExtraParams: extraParams,
			Conn:        conn,
		}

		if err != nil {
			ReleaseClient(conn, message.GenerateGenericErrorHttpResponse())
		}

		if operation == core.OpLock ||
			operation == core.OpUnlock {

			app.ChanLocks <- operationData
		} else if operation == core.OpReserveStock ||
			operation == core.OpReleaseStock ||
			operation == core.OpCreateStock {

			app.ChanStocks <- operationData
		} else if operation == core.OpSetupRateLimiter ||
			operation == core.OpWaitForRateLimiter {

			app.ChanRl <- operationData
		}
	}
}
