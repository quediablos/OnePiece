package core

import (
	"net"
	"time"
)

type LockInfo struct {
	ResourceId string
	IsLocked   bool
	LockedAt   *time.Time
}

type LockRequest struct {
	ResourceId  string
	RequestedBy net.Conn
	RequestedAt *time.Time
}

func CheckForLock(resourceId string, Locks map[string]LockInfo) *LockInfo {

	if lock, ok := Locks[resourceId]; ok {
		return &lock
	} else {
		return nil
	}
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
