package main

import (
	"OnePiece/core"
	"OnePiece/network"
)

func main() {

	//Start the app data.
	app := core.NewApp()

	network.ListenHttp(app)
}
