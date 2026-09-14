package main

import (
	"OnePiece/core"
	"OnePiece/job"
	"OnePiece/network"
)

func main() {

	//Start the app data.
	app := core.NewApp()

	//Maintenance jobs
	go job.Maintain(app)

	//Workers
	go network.ProcessLocks(app)
	go network.ProcessStocks(app)

	//Listener
	network.ListenHttp(app)

}
