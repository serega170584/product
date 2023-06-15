package main

import "fmt"

func main() {
	ch1 := make(chan int, 20)
	ch2 := make(chan int, 20)
	ch3 := make(chan int, 20)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	ch2 <- 4
	ch2 <- 5
	ch2 <- 6
	ch3 <- 7
	ch3 <- 8
	close(ch1)
	close(ch2)
	close(ch3)

	chRes := mergeChannels(ch1, ch2, ch3)
	for num := range chRes {
		fmt.Println(num)
	}
}

func mergeChannels(chs ...chan int) chan int {
	res := make(chan int)
	chCounter := make(chan int)

	go func() {
		chCounter <- len(chs)
	}()

	for _, vals := range chs {
		vals := vals
		go func() {
			for val := range vals {
				res <- val
			}
			cnt := <-chCounter
			cnt--
			if cnt == 0 {
				close(res)
			}
			chCounter <- cnt
		}()
	}
	return res
}
