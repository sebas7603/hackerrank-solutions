package main

import "fmt"

// Birthday Cake Candles
// Problem: https://www.hackerrank.com/challenges/birthday-cake-candles/problem
// Author:  Sebastián Ayala Suárez
// Since:   29/07/2023

func birthdayCakeCandles(candles []int32) int32 {
	tallest := candles[0]
	quantity := 1
	for i := 1; i < len(candles); i++ {
		if candles[i] > tallest {
			tallest = candles[i]
			quantity = 1
		} else if candles[i] == tallest {
			quantity++
		}
	}
	return int32(quantity)
}

// The following code is not part of the solution but is necessary for local functionality.

func main() {
	arr := []int32{3, 2, 1, 3}
	res := birthdayCakeCandles(arr)
	fmt.Printf("%d\n", res)
	return
}
