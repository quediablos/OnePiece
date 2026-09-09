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
		if lock.LockedAt != nil && time.Since(*lock.LockedAt) > 10*time.Second {
			lock.Expired = true
			delete(locks, resourceId)
		}

		return &lock
	}

	return nil
}

func AcquireLock(resourceId string, Locks map[string]LockInfo) {

	now := time.Now()
	Locks[resourceId] = LockInfo{
		ResourceId: resourceId,
		IsLocked:   true,
		LockedAt:   &now,
	}
}

func RequestLock(Conn net.Conn, resourceId string, LockRequests map[string][]LockRequest) {

	if LockRequests[resourceId] == nil {
		LockRequests[resourceId] = []LockRequest{}
	}

	now := time.Now()
	LockRequests[resourceId] = append(LockRequests[resourceId], LockRequest{
		ResourceId:  resourceId,
		RequestedBy: Conn,
		RequestedAt: &now,
	})
}

func ReleaseLock(resourceId string, Locks map[string]LockInfo, lockRequests map[string][]LockRequest) net.Conn {

	delete(Locks, resourceId)

	if queue, ok := lockRequests[resourceId]; ok && len(queue) > 0 {
		waitingClient := queue[0]
		lockRequests[resourceId] = queue[1:]
		AcquireLock(resourceId, Locks)

		return waitingClient.RequestedBy
	}
	return nil
}
