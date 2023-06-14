package main

import (
	"fmt"
	"sync"
)

func main() {
	writes := 1000
	m := make(map[int]int, writes)
	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}
	wg.Add(writes)
	for i := 0; i < writes; i++ {
		i := i
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}()
	}

	reads := 1000
	wg.Add(reads)
	for i := 0; i < reads; i++ {
		i := i
		go func() {
			defer wg.Done()
			mu.RLock()
			_, _ = m[i]
			mu.RUnlock()
		}()
	}

	wg.Wait()

	fmt.Println(m)
}
