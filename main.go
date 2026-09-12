package main

import (
	"OnePiece/core"
	"OnePiece/job"
	"OnePiece/network"
)

func main() {

	//Start the app data.
	app := core.NewApp()

	go job.Maintain(app)

	network.ListenForLocksHttp(app)
}
