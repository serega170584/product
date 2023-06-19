package main

import (
	"fmt"
	"sync"
)

func main() {
	cnt := 10
	m := make(map[int]int, cnt)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		i := i
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			m[i] = i
		}()
	}

	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		i := i
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			fmt.Println(m[i])
		}()
	}

	wg.Wait()

}
