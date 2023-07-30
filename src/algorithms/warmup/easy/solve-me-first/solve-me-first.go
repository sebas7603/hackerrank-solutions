package main

import "fmt"

// Solve Me First
// Problem: https://www.hackerrank.com/challenges/solve-me-first/problem
// Author:  Sebastián Ayala Suárez
// Since:   29/07/2023

func solveMeFirst(a uint32, b uint32) uint32 {
	return a + b
}

// The following code is not part of the solution but is necessary for local functionality.

func main() {
	var a, b, res uint32
	a = 5
	b = 18
	res = solveMeFirst(a, b)
	fmt.Println(res)
}
