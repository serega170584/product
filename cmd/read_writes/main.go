package main

import (
	"fmt"
	"sync"
)

func main() {
	writes := 1000
	m := make(map[int]int, writes)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
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
			mu.Lock()
			_, _ = m[i]
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(m)
}
