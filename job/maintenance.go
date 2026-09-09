package job

import (
	"OnePiece/core"
	"OnePiece/message"
	"OnePiece/network"
	"time"
)

func Maintain(app *core.App) {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		MaintainPendingLockRequests(app)
	}
}

func MaintainPendingLockRequests(app *core.App) {

	app.MutexForLocks.Lock()
	defer app.MutexForLocks.Unlock()

	_, waitingLockRequests := core.CheckForExpiredLocks(app)

	for _, req := range waitingLockRequests {
		core.AcquireLock(req.ResourceId, app)
		network.ReleaseClientHttp(req.RequestedBy, message.GenerateAcquireLockResponse(req.ResourceId))
	}
}
