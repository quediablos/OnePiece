package core

import "time"

type LockInfo struct {

	ResourceId string
	IsLocked   *bool
	LockedAt   *time.Time
}

func CheckForLock(resourceId string, Locks map[string]LockInfo) *LockInfo {

	if lock, ok := Locks[resourceId]; !ok {
		return &lock
	} else {
		return nil
	}
}