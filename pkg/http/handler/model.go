package handler

import "sync"

type IHandler interface {
	GetPlayers() *sync.Map
}
