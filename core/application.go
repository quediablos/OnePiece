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

	//Rate limiter
	RateLimiters map[string]*RateLimiter //Key is resourceId:userId

	//Thread-safe
	MutexForLocks  sync.RWMutex
	MutexForStocks sync.RWMutex
	MutexForRl     sync.RWMutex

	//Thread communication
	ChanLocks  chan OperationData
	ChanStocks chan OperationData
	ChanRl     chan OperationData

	//Config
	Config Config
}

type Config struct {
	ListenHttp bool
}

func NewApp() *App {
	return &App{
		Locks:         make(map[string]LockInfo),
		LockRequests:  make(map[string][]LockRequest),
		StockCounts:   make(map[string]int64),
		StockReserves: make(map[string][]StockReserve),
		RateLimiters:  make(map[string]*RateLimiter),
		ChanLocks:     make(chan OperationData, 100),
		ChanStocks:    make(chan OperationData, 100),
		ChanRl:        make(chan OperationData, 100),
		Config:        Config{ListenHttp: true},
	}
}

func (app *App) IncrementCycle() int32 {
	app.Cycle++
	return app.Cycle
}
