package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 6)
	ch3 := make(chan int, 10)

	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	ch2 <- 4
	ch2 <- 5
	ch2 <- 6
	ch2 <- 4
	ch2 <- 5
	ch2 <- 6
	ch3 <- 7
	ch3 <- 8
	ch3 <- 9
	ch3 <- 10
	ch3 <- 11
	ch3 <- 12
	ch3 <- 13
	ch3 <- 14
	ch3 <- 15
	ch3 <- 16

	close(ch1)
	close(ch2)
	close(ch3)

	ch := test(ch1, ch2, ch3)

	for val := range ch {
		fmt.Printf("Elem: %d\n", val)
	}
}

func test(chList ...chan int) chan int {
	res := make(chan int)
	counterCh := make(chan struct{})
	var cnt int

	for _, chVals := range chList {
		chVals := chVals
		cnt = cnt + len(chVals)

		go func() {
			for val := range chVals {
				res <- val
				<-counterCh
			}
		}()
	}

	go func() {
		for cnt > 0 {
			counterCh <- struct{}{}
			cnt--
		}

		close(res)
	}()

	return res
}
