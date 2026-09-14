package core

import (
	"net"
	"time"
)

// LockInfo Operations in core level handle the data level only.
type LockInfo struct {
	ResourceId string
	IsLocked   bool
	LockedAt   *time.Time
	Expired    bool
}

type LockRequest struct {
	ResourceId  string
	RequestedBy net.Conn
	RequestedAt *time.Time
}

// CheckForLock Checks if the resourceId is validly under lock. If the locking client has not locked before the TTL
func CheckForLock(resourceId string, locks map[string]LockInfo) *LockInfo {

	if lock, ok := locks[resourceId]; ok {

		//Check for expiration. If the previous lock expired, delete it from the locks.
		if lock.LockedAt != nil && time.Since(*lock.LockedAt) > lockExpirationTtl {
			lock.Expired = true
			delete(locks, resourceId)
		}

		return &lock
	}

	return nil
}

// AcquireLock
// ------------- THREAD-SAFE: This method needs to run thread-safe -------------
func AcquireLock(resourceId string, app *App) {

	now := time.Now()
	app.Locks[resourceId] = LockInfo{
		ResourceId: resourceId,
		IsLocked:   true,
		LockedAt:   &now,
	}
}

// RequestLock
// ------------- THREAD-SAFE: This method needs to run thread-safe -------------
func RequestLock(Conn net.Conn, resourceId string, app *App) {

	if app.LockRequests[resourceId] == nil {
		app.LockRequests[resourceId] = []LockRequest{}
	}

	now := time.Now()
	app.LockRequests[resourceId] = append(app.LockRequests[resourceId], LockRequest{
		ResourceId:  resourceId,
		RequestedBy: Conn,
		RequestedAt: &now,
	})
}

func ReleaseLock(resourceId string, app *App) net.Conn {

	delete(app.Locks, resourceId)

	if queue, ok := app.LockRequests[resourceId]; ok && len(queue) > 0 {
		waitingClient := queue[0]
		app.LockRequests[resourceId] = queue[1:]
		AcquireLock(resourceId, app)

		return waitingClient.RequestedBy
	}
	return nil
}

// CheckForExpiredLocks checks the locks map for any LockInfo whose LockedAt is
// older than lockExpirationTtl and removes them.
// ------------- THREAD-SAFE: This method needs to run thread-safe -------------
func CheckForExpiredLocks(app *App) ([]LockInfo, []LockRequest) {

	var expiredLocks []LockInfo
	var pendingRequests []LockRequest
	for resourceId, lock := range app.Locks {
		if lock.LockedAt != nil && time.Since(*lock.LockedAt) > lockExpirationTtl {
			lock.Expired = true
			expiredLocks = append(expiredLocks, lock)
			delete(app.Locks, resourceId)

			if queue, ok := app.LockRequests[resourceId]; ok && len(queue) > 0 {
				pendingRequests = append(pendingRequests, queue[0])
				app.LockRequests[resourceId] = queue[1:]
			}
		}
	}
	return expiredLocks, pendingRequests
}
