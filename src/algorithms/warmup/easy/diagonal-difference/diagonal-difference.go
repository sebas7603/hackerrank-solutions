package main

import (
	"fmt"
)

// Diagonal Difference
// Problem: https://www.hackerrank.com/challenges/diagonal-difference/problem
// Author:  Sebastián Ayala Suárez
// Since:   29/07/2023
// Note:    recursiveDiag is a recursive function.

func recursiveDiag(arr [][]int32, i int) (int32, int32) {
	if i == len(arr[0]) {
		return 0, 0
	}
	val_1, val_2 := recursiveDiag(arr, i+1)
	return arr[i][i] + val_1, arr[i][len(arr[0])-1-i] + val_2
}

func diagonalDifference(arr [][]int32) (res int32) {
	diag_1, diag_2 := recursiveDiag(arr, 0)
	if diag_1 > diag_2 {
		res = diag_1 - diag_2
		return
	}
	res = diag_2 - diag_1
	return
}

// The following code is not part of the solution but is necessary for local functionality.

func main() {
	matrix := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	result := diagonalDifference(matrix)
	fmt.Printf("%d\n", result)
	return
}
