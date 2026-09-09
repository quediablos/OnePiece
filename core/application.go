package core

import (
	"sync"
	"time"
)

const (
	lockExpirationTtl = 10 * time.Second
)

type App struct {
	Locks         map[string]LockInfo      //Key is the resource id.
	LockRequests  map[string][]LockRequest //Key is the resource id.
	Cycle         int32                    //Counts each operation handled, in some cycles maintenance work is done.
	MutexForLocks sync.RWMutex             //Guards Locks and LockRequests.
}

func NewApp() *App {
	return &App{
		Locks:        make(map[string]LockInfo),
		LockRequests: make(map[string][]LockRequest),
	}
}

func (app *App) IncrementCycle() int32 {
	app.Cycle++
	return app.Cycle
}
