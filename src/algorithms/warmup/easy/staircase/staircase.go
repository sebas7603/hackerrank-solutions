package main

import (
	"fmt"
	"strings"
)

// Staircase
// Problem: https://www.hackerrank.com/challenges/staircase/problem
// Author:  Sebastián Ayala Suárez
// Since:   29/07/2023
// Note:    printStep is a recursive function.

func printStep(n int32, step int32) int32 {
	fmt.Printf("%s%s\n", strings.Repeat(" ", int(n-step)), strings.Repeat("#", int(step)))
	if step == n {
		return 0
	}
	return printStep(n, step+1)
}

func staircase(n int32) {
	printStep(n, 1)
	return
}

// The following code is not part of the solution but is necessary for local functionality.

func main() {
	staircase(30)
	return
}
