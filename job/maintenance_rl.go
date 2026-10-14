package job

import (
	"OnePiece/core"
	"OnePiece/message"
	"OnePiece/network"
	"container/heap"
	"fmt"
	"time"
)

func MaintainRateLimiter(app *core.App) {
	ticker := time.NewTicker(1000 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		HandleQueuedClients(app)
	}
}

// HandleQueuedClients pops and prints all waiting clients from the global priority queue.
func HandleQueuedClients(app *core.App) {
	app.MutexForRl.Lock()
	defer app.MutexForRl.Unlock()

	fmt.Printf("[RateLimiter] Waiting clients (%d):\n", app.WaitingClientsAll.Len())
	for i := 0; app.WaitingClientsAll.Len() > 0; i++ {

		client := heap.Pop(&app.WaitingClientsAll).(core.WaitingClient)

		released, _, _ := core.CheckAndMaintainRlAvailability(app, client.Key, client.OperationData, true)

		if released {
			fmt.Printf("[RateLimiter] Client (%d) is now released.\n", client.Key)
			network.ReleaseClient(client.OperationData.Conn, message.GenerateWaitRlSuccessfulResponse(app))
		} else {
			fmt.Printf("[RateLimiter] Client (%d) keeps waiting.\n", client.Key)
		}

	}
}
