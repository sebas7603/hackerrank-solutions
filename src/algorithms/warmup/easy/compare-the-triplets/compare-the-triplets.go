package main

import "fmt"

// Compare The Triplets
// Problem: https://www.hackerrank.com/challenges/compare-the-triplets/problem
// Author:  Sebastián Ayala Suárez
// Since:   29/07/2023

func compareTriplets(a []int32, b []int32) (rating [2]int32) {
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			rating[0]++
			continue
		} else if b[i] > a[i] {
			rating[1]++
		}
	}
	return
}

// The following code is not part of the solution but is necessary for local functionality.

func printResult(result [2]int32) {
	for i, resultItem := range result {
		fmt.Printf("%d", resultItem)

		if i != len(result)-1 {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
}

func main() {
	a := []int32{5, 6, 7}
	b := []int32{3, 6, 10}
	result := compareTriplets(a, b)
	printResult(result)

	a2 := []int32{17, 28, 30}
	b2 := []int32{99, 16, 8}
	result = compareTriplets(a2, b2)
	printResult(result)
	return
}
