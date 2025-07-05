package aiclient

import (
	"sync"
)

type ClientManager struct {
	clients map[string]AIClient
	mutex   sync.RWMutex
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		clients: make(map[string]AIClient),
	}
}
