package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 1000
	m := make(map[int]int, count)

	wg := sync.WaitGroup{}
	wg.Add(count)
	mu := sync.Mutex{}
	for i := 0; i < count; i++ {
		i := i
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(m)
}
