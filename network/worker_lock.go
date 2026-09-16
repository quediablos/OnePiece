package network

import (
	"OnePiece/core"
	"OnePiece/message"
)

func ProcessLocks(app *core.App) {
	for data := range app.ChanLocks {

		app.MutexForLocks.Lock()

		if data.Operation == core.OpLock {

			lock := core.CheckForLock(data.ResourceId, app.Locks)

			if lock != nil && !lock.Expired {
				// Another client already holds the lock — queue this one.
				core.RequestLock(data.Conn, data.ResourceId, app)
			} else {
				// No active lock — acquire it and respond immediately.
				core.AcquireLock(data.ResourceId, app)
				ReleaseClient(data.Conn, message.GenerateAcquireLockResponse(app, data.ResourceId))
			}

		} else if data.Operation == core.OpUnlock {

			waitingOne := core.ReleaseLock(data.ResourceId, app)

			if waitingOne != nil {
				ReleaseClient(waitingOne, message.GenerateAcquireLockResponse(app, data.ResourceId))
			}
			ReleaseClient(data.Conn, message.GenerateReleaseLockResponse(app, data.ResourceId))
		}

		app.MutexForLocks.Unlock()
	}
}
