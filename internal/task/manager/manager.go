package manager

import (
	"fmt"
	"time"
)

type Manager struct {
	ch    chan bool
	count int
}

func NewManager(mainChannel chan bool, cnt int) *Manager {
	for i := 0; i < cnt; i++ {
		mainChannel <- false
	}
	return &Manager{count: cnt, ch: mainChannel}
}

func (manager *Manager) Execute() error {
	for i := 0; i < manager.count; i++ {
		go func(i int) {
			time.Sleep(time.Duration(10 * time.Second))
			fmt.Printf("Task number %d\n", i)
			<-manager.ch
		}(i)
	}
	return nil
}
