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

	var notReleased []core.EnqueuedClient

	for app.EnqueuedClientsAll.Len() > 0 {

		client := heap.Pop(&app.EnqueuedClientsAll).(core.EnqueuedClient)
		released, _ := core.CheckAndMaintainRlAvailabilityForEnqueuedClient(app, client.Key, client)

		if released {
			fmt.Printf("[RateLimiter] Client (key:%s priority:%f) is now released.\n", client.Key, client.Priority)
			network.ReleaseClient(client.OperationData.Conn, message.GenerateWaitRlSuccessfulResponse(app))
		} else {
			notReleased = append(notReleased, client)
		}
	}

	for _, client := range notReleased {
		heap.Push(&app.EnqueuedClientsAll, client)
	}
}
