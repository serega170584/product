package context

import (
	"context"
	"fmt"
	"time"
)

type Manager struct {
	count int
}

func New(c int) *Manager {
	return &Manager{count: c}
}

func (manager *Manager) Execute(ctx context.Context) {
	for i := 0; i < manager.count; i++ {
		go doneContext(ctx)
	}
}

func doneContext(ctx context.Context) {
	defer func() {
		fmt.Println("Goroutine is worked")
	}()
loop:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context done")
			break loop
		default:
			<-time.After(3 * time.Second)
			fmt.Println("context in work")
		}
	}
}
