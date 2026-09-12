package core

import (
	"sync"
	"time"
)

const (
	lockExpirationTtl = 10 * time.Second
)

type App struct {
	Cycle int32 //Counts each operation handled, in some cycles maintenance work is done.

	//Locks
	Locks        map[string]LockInfo      //Key is the resource id.
	LockRequests map[string][]LockRequest //Key is the resource id.

	//Stocks
	StockCounts   map[string]int64          //Represents the stock count for each resource.
	StockReserves map[string][]StockReserve //Stores the stock reserves for each stock.

	//Thread-safe
	MutexForLocks  sync.RWMutex //Guards Locks and LockRequests.
	MutexForStocks sync.RWMutex
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
