package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	count := 1000
	doubles := make([]int, count)
	existed := make(map[int]struct{}, count)
	unique := make(chan int, count)

	for i := 0; i < count; i++ {
		doubles[i] = rand.Intn(10)
	}

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		i := i
		go func() {
			defer wg.Done()
			mu.Lock()
			if _, ok := existed[doubles[i]]; !ok {
				existed[doubles[i]] = struct{}{}
				unique <- doubles[i]
			}
			mu.Unlock()
		}()
	}
	wg.Wait()

	close(unique)
	for val := range unique {
		fmt.Println(val)
	}
}
