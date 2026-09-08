package main

import "OnePiece/core"

type App struct {
	Locks        map[string]core.LockInfo
	LockRequests map[string][]core.LockRequest //Key is the resource id.
}

func NewApp() *App {
	return &App{
		Locks:        make(map[string]core.LockInfo),
		LockRequests: make(map[string][]core.LockRequest),
	}
}
