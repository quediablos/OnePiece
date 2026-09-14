package network

import (
	"OnePiece/core"
	"OnePiece/message"
	"fmt"
	"log"
	"net"
	"strconv"
)

func ListenForStocksHttp(app *core.App) {

	listener, err := net.Listen("tcp", "127.0.0.1:3004")
	if err != nil {
		log.Fatalf("Failed to bind to port: %v", err)
	}
	defer listener.Close()
	fmt.Println("Server running on http://127.0.0.1:3004...")

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
		handleOperationStock(operation, resourceId, extraParams, conn, app)
	}
}

func handleOperationStock(operation core.Operation, resourceId string, extraParams []string, conn net.Conn, app *core.App) {

	app.MutexForStocks.Lock()
	defer app.MutexForStocks.Unlock()

	if operation == core.OpReserveStock {

		stockReserve, success := app.ReserveStock(resourceId)

		if success {
			ReleaseClientHttp(conn, message.GenerateReserveStockSuccessfulResponse(stockReserve.Id, resourceId))
		} else {
			ReleaseClientHttp(conn, message.GenerateReserveStockFailedResponse(resourceId))
		}

	} else if operation == core.OpReleaseStock {

	} else if operation == core.OpCreateStock {

		quantity, _ := strconv.ParseInt(extraParams[0], 10, 64)
		app.CreateStock(resourceId, quantity)

		ReleaseClientHttp(conn, message.GenerateCreateStockResponse(resourceId, extraParams[0]))
	}
}
