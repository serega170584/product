package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 20
	wg := sync.WaitGroup{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		i := i
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}
