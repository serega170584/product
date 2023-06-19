package main

import "fmt"

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 123
	ch1 <- 456
	ch1 <- 789
	ch2 := make(chan int, 3)
	ch2 <- 1234
	ch2 <- 5678
	ch2 <- 9012
	ch3 := make(chan int, 3)
	ch3 <- 12340
	ch3 <- 67890
	ch3 <- 12345

	ch := mergeChannels(ch1, ch2, ch3)

	for val := range ch {
		fmt.Println(val)
	}
}

func mergeChannels(chList ...chan int) chan int {
	length := 0
	counter := make(chan int)
	for _, ch := range chList {
		length = length + len(ch)
	}

	res := make(chan int, length)

	for _, chVals := range chList {
		chVals := chVals
		go func() {
			for val := range chVals {
				res <- val
				length--
				counter <- length
			}
		}()
	}

	go func() {
		for cnt := range counter {
			if cnt == 0 {
				close(res)
			}
		}
	}()

	return res
}
