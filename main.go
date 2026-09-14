package main

import (
	"OnePiece/core"
	"OnePiece/job"
	"OnePiece/network"
	"sync"
)

func main() {

	//Start the app data.
	app := core.NewApp()

	go job.Maintain(app)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		network.ListenForLocksHttp(app)
	}()

	go func() {
		defer wg.Done()
		network.ListenForStocksHttp(app)
	}()

	wg.Wait()
}
