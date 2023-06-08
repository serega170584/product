package manager

import (
	"fmt"
	"sync"
	"time"
)

type Manager struct {
	count     int
	waitGroup *sync.WaitGroup
}

func NewManager(cnt int, wg *sync.WaitGroup) *Manager {
	return &Manager{count: cnt, waitGroup: wg}
}

func (manager *Manager) Execute() error {
	routine := func(i int) {
		defer manager.waitGroup.Done()
		ticker := time.NewTicker(10 * time.Second)
		func() {
			<-ticker.C
			fmt.Printf("Task number %d\n", i)
			ticker.Stop()
		}()
	}

	for i := 0; i < manager.count; i++ {
		go routine(i)
	}
	return nil
}
