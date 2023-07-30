package main

import "fmt"

// A Very Big Sum
// Problem: https://www.hackerrank.com/challenges/a-very-big-sum/problem
// Author:  Sebastián Ayala Suárez
// Since:   29/07/2023
// Note:    recursiveSum is a recursive function.

func recursiveSum(ar []int64, i int) int64 {
	if i == len(ar) {
		return 0
	}
	return ar[i] + recursiveSum(ar, i+1)
}

func aVeryBigSum(ar []int64) (res int64) {
	return recursiveSum(ar, 0)
}

// The following code is not part of the solution but is necessary for local functionality.

func main() {
	arr := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	result := aVeryBigSum(arr)
	fmt.Printf("%d\n", result)
	return
}
