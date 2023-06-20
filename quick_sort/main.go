package main

import (
	"fmt"
	"math/rand"
)

func main() {
	cnt := 20
	leftRightValues := make([]leftRightVal, 0)
	list := make([]int, cnt)
	for i := 0; i < 20; i++ {
		list[i] = rand.Intn(30) + 10
	}
	list = []int{10, 11, 27, 11, 31, 27, 22, 32, 16, 14, 34, 31, 30, 22, 35, 17, 11, 28, 33, 28}

	leftRightValues = append(leftRightValues, leftRightVal{left: 1, right: cnt - 1})
	counter := 1
	currIndex := 0
	fmt.Println(list)
loop:
	for counter > 0 {
		counter--
		val := leftRightValues[currIndex]
		leftInd := val.left
		rightInd := val.right
		if leftInd >= rightInd {
			currIndex++
			continue loop
		}
		num := list[leftInd]
		fixInd := leftInd
		leftInd++
		for leftInd <= rightInd {
			for leftInd <= val.right && list[leftInd] < num {
				leftInd++
			}

			for rightInd >= val.left && list[rightInd] > num {
				rightInd--
			}

			if leftInd <= rightInd {
				list[leftInd], list[rightInd] = list[rightInd], list[leftInd]
				leftInd++
				rightInd--
			}
		}

		list[fixInd], list[rightInd] = list[rightInd], list[fixInd]

		leftRightValues = append(leftRightValues, leftRightVal{left: fixInd, right: rightInd - 1})
		counter++

		leftRightValues = append(leftRightValues, leftRightVal{left: rightInd + 1, right: val.right})
		counter++

		currIndex++
	}

	fmt.Println(list)
}

type leftRightVal struct {
	left  int
	right int
}
