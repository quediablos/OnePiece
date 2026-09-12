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

		operation, resourceId, err := req.ParseURL()
		handleOperation(operation, resourceId, conn, app)
	}
}

func handleOperation(operation core.Operation, resourceId string, conn net.Conn, app *core.App) {

	app.MutexForLocks.Lock()
	defer app.MutexForLocks.Unlock()

	if operation == core.Lock {

		lock := core.CheckForLock(resourceId, app.Locks)

		if lock != nil && !lock.Expired {
			//Another client already locked the resource, hold this client until the lock is released.
			core.RequestLock(conn, resourceId, app)
		} else {

			//No lock is acquired for the resource id yet, acquire the lock, and finish the connectin.
			core.AcquireLock(resourceId, app)
			ReleaseClientHttp(conn, message.GenerateAcquireLockResponse(resourceId))
		}
	} else if operation == core.Unlock {

		waitingOne := core.ReleaseLock(resourceId, app)

		if waitingOne != nil {
			ReleaseClientHttp(waitingOne, message.GenerateAcquireLockResponse(resourceId))
		}
		ReleaseClientHttp(conn, message.GenerateReleaseLockResponse(resourceId))
	}

}
