package main

import "fmt"

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	ch2 := make(chan int, 3)
	ch2 <- 4
	ch2 <- 5
	ch2 <- 6
	ch3 := make(chan int, 3)
	ch3 <- 7
	ch3 <- 8
	ch3 <- 9

	ch := mergeChannels(ch1, ch2, ch3)
	for val := range ch {
		fmt.Println(val)
	}
}

func mergeChannels(chList ...chan int) chan int {
	res := make(chan int)
	counter := make(chan struct{})
	cnt := 0

	for _, chVals := range chList {
		cnt = cnt + len(chVals)
		chVals := chVals
		go func() {
			for chVal := range chVals {
				res <- chVal
				counter <- struct{}{}
			}
		}()
	}

	go func() {
		for range counter {
			cnt--
			if cnt == 0 {
				close(res)
			}
		}
	}()

	return res
}
