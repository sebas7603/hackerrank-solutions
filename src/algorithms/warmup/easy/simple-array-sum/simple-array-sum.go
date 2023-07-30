package main

import "fmt"

// Simple Array Sum
// Problem: https://www.hackerrank.com/challenges/simple-array-sum/problem
// Author:  Sebastián Ayala Suárez
// Since:   29/07/2023

func simpleArraySum(ar []int32) int32 {
	sum := int32(0)
	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}
	return sum
}

// The following code is not part of the solution but is necessary for local functionality.

func main() {
	arr := []int32{1, 2, 3, 4, 10, 11}
	result := simpleArraySum(arr)
	fmt.Printf("%d\n", result)

	arr = []int32{5, 5, 5}
	result = simpleArraySum(arr)
	fmt.Printf("%d\n", result)
}
