package network

import (
	"OnePiece/core"
	"OnePiece/message"
	"fmt"
	"log"
	"net"
	"strings"
)

const (
	TCP_MESSAGE_BUFFER_SIZE = 256
)

// ListenTcp
// Message frames are referred in README.md.
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

		buf := make([]byte, TCP_MESSAGE_BUFFER_SIZE)
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("Failed to read from connection: %v", err)
			conn.Close()
			continue
		}

		input := buf[:n]
		operationData, err := parseTcpInputV1(conn, string(input))

		if err != nil {
			log.Printf("Failed to parse input: %v", err)
			ReleaseClient(conn, message.GenerateErrorHttpResponse(err.Error()))
			continue
		}

		if operationData.Operation == core.OpLock || operationData.Operation == core.OpUnlock {
			app.ChanLocks <- operationData

		} else if operationData.Operation == core.OpCreateStock ||
			operationData.Operation == core.OpReserveStock ||
			operationData.Operation == core.OpReleaseStock {

			app.ChanStocks <- operationData
		}
	}

}

// parseTcpInputV1 parses the raw tcp input to return operation data.
func parseTcpInputV1(conn net.Conn, input string) (core.OperationData, error) {
	input = strings.TrimRight(input, "\n\r")
	parts := strings.Split(input, "|")
	if len(parts) < 4 {
		return core.OperationData{},
			fmt.Errorf("invalid message format: expected at least 4 pipe-separated fields, got %d", len(parts))
	}

	// parts[0] = version, parts[1] = messageType, parts[2] = operation, parts[3] = resourceId
	opRaw := strings.ToUpper(strings.TrimSpace(parts[2]))
	resourceId := strings.TrimSpace(parts[3])

	var op core.Operation
	switch opRaw {
	case "LOCK":
		op = core.OpLock
	case "UNLOCK":
		op = core.OpUnlock
	case "CREATE_STOCK":
		op = core.OpCreateStock
	case "RESERVE_STOCK":
		op = core.OpReserveStock
	case "RELEASE_STOCK":
		op = core.OpReleaseStock
	default:
		return core.OperationData{}, fmt.Errorf("unknown operation: %q", opRaw)
	}

	var extraParams []string
	if len(parts) > 4 {
		extraParams = parts[4:]
	}

	return core.OperationData{
		Operation:   op,
		ResourceId:  resourceId,
		ExtraParams: extraParams,
		Conn:        conn,
	}, nil
}
