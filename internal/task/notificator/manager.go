package notificator

import (
	"fmt"
	"time"
)

type Manager struct {
	count int
}

func New(cnt int) *Manager {
	return &Manager{count: cnt}
}

func (manager *Manager) Execute(ch <-chan struct{}) {
	for i := 0; i < manager.count; i++ {
		go doneNotification(ch)
	}
}

func doneNotification(ch <-chan struct{}) {
	defer func() {
		fmt.Println("Goroutine is worked")
	}()
loop:
	for {
		select {
		case <-ch:
			fmt.Println("Channel")
			break loop
		default:
			fmt.Println("Goroutine in work")
			<-time.After(3 * time.Second)
		}
	}
}
