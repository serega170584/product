package main

import (
	"fmt"
	"sync"
)

func main() {
	cnt := 10
	m := make(map[int]int, cnt)
	mu := &sync.Mutex{}

	counter := NewCounter(mu)
	counter.addCount(cnt)

	for i := 0; i < cnt; i++ {
		i := i
		go func() {
			defer counter.done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}()
	}

	counter.addCount(cnt)

	for i := 0; i < cnt; i++ {
		i := i
		go func() {
			defer counter.done()
			mu.Lock()
			fmt.Printf("Elem: %d\n", m[i])
			mu.Unlock()
		}()
	}

	counter.wait()
}

type Counter struct {
	cnt int
	mu  *sync.Mutex
}

func NewCounter(mu *sync.Mutex) *Counter {
	counter := &Counter{}
	counter.mu = mu

	return counter
}

func (counter *Counter) addCount(cnt int) {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	counter.cnt = cnt + counter.cnt
}

func (counter *Counter) done() {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	counter.cnt--
}

func (counter *Counter) wait() {
loop:
	for {
		if counter.cnt == 0 {
			break loop
		}
	}
}
