package main

import "OnePiece/core"

type App struct {
	Locks map[string]core.LockInfo
}

func NewApp() *App {
	return &App{
		Locks: make(map[string]core.LockInfo),
	}
}
