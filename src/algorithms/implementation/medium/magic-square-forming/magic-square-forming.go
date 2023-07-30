package main

import (
	"fmt"
	"math"
)

// Forming a Magic Square
// Problem: https://www.hackerrank.com/challenges/magic-square-forming/problem
// Author:  Sebastián Ayala Suárez
// Since:   29/07/2023
// Note:    The value of magicSquares contains all possible combinations for a 3x3 magic square.

func formingMagicSquare(s [][]int32) (res int32) {
	var magicSquares = [8][3][3]int32{
		{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
		{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
		{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
		{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
		{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
		{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
		{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
		{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
	}
	var results [8]int32
	for i := range magicSquares {
		for j := range magicSquares[i] {
			for k := range magicSquares[i][j] {
				results[i] += int32(math.Abs(float64(magicSquares[i][j][k] - s[j][k])))
			}
		}
	}

	res = results[0]
	for x := 1; x < 8; x++ {
		if results[x] < res {
			res = results[x]
		}
	}

	return
}

// The following code is not part of the solution but is necessary for local functionality.

func main() {
	matrix := [][]int32{
		{4, 9, 2},
		{3, 5, 7},
		{8, 1, 5},
	}
	result := formingMagicSquare(matrix)
	fmt.Printf("%d\n", result)

	matrix = [][]int32{
		{4, 8, 2},
		{4, 5, 7},
		{6, 1, 6},
	}
	result = formingMagicSquare(matrix)
	fmt.Printf("%d\n", result)
	return
}
