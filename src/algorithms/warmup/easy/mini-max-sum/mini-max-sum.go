package main

import (
	"fmt"
)

// mini-max-sum
// Problem: https://www.hackerrank.com/challenges/mini-max-sum/problem
// Author:  Sebastián Ayala Suárez
// Since:   29/07/2023
// Note:    To solve this problem, calculate the sum of all array elements and then
//          subtract the minimum value to obtain maxSum, and subtract the maximum value to obtain minSum.

func miniMaxSum(arr []int32) {
	var (
		min int32 = arr[0]
		max int32 = arr[0]
		sum int64
	)

	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		sum += int64(num)
	}

	fmt.Printf("%d %d\n", sum-int64(max), sum-int64(min))
}

// The following code is not part of the solution but is necessary for local functionality.

func main() {
	arr := []int32{1, 2, 3, 4, 5}
	miniMaxSum(arr)

	arr2 := []int32{7, 69, 2, 221, 8974}
	miniMaxSum(arr2)

	arr3 := []int32{793810624, 895642170, 685903712, 623789054, 468592370}
	miniMaxSum(arr3)

	return
}
