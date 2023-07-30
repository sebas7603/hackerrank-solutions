package main

import (
	"fmt"
)

// Plus Minus
// Problem: https://www.hackerrank.com/challenges/plus-minus/problem
// Author:  Sebastián Ayala Suárez
// Since:   29/07/2023
// Note:    getRatios is called in a gorroutine, I use channels to get the results

func getRatios(arr []int32, resChan chan<- float64) {
	positiveCount := 0
	negativeCount := 0
	zeroCount := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			positiveCount++
			continue
		} else if arr[i] < 0 {
			negativeCount++
			continue
		}
		zeroCount++
	}

	resChan <- float64(positiveCount) / float64(len(arr))
	resChan <- float64(negativeCount) / float64(len(arr))
	resChan <- float64(zeroCount) / float64(len(arr))
	close(resChan)
}

func plusMinus(arr []int32) {
	resChan := make(chan float64, 3)
	go getRatios(arr, resChan)

	fmt.Printf("%.6f\n", <-resChan)
	fmt.Printf("%.6f\n", <-resChan)
	fmt.Printf("%.6f\n", <-resChan)
}

// The following code is not part of the solution but is necessary for local functionality.

func main() {
	arr := []int32{-4, 3, -9, 0, 4, 1}
	plusMinus(arr)

	arr2 := []int32{1, 2, 3, -1, -2, -3, 0, 0}
	plusMinus(arr2)

	return
}
