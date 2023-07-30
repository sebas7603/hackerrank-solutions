package main

import (
	"fmt"
)

// Climbing the Leaderboard
// Problem: https://www.hackerrank.com/challenges/climbing-the-leaderboard/problem
// Author:  Sebastián Ayala Suárez
// Since:   30/07/2023
// Note:    getPositionByBinarySearch uses binary search to find the position of a score

func getPositionByBinarySearch(uniqueRankeds []int32, playerScore int32) int {
	first, last := 0, len(uniqueRankeds)-1
	for first <= last {
		mid := (first + last) / 2
		if uniqueRankeds[mid] == playerScore {
			return mid
		} else if uniqueRankeds[mid] < playerScore {
			last = mid - 1
		} else {
			first = mid + 1
		}
	}
	return first
}

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	uniqueRankeds := make([]int32, 0, len(player))
	uniqueRankeds = append(uniqueRankeds, ranked[0])
	for i := 1; i < len(ranked); i++ {
		if ranked[i] != ranked[i-1] {
			uniqueRankeds = append(uniqueRankeds, ranked[i])
		}
	}

	results := make([]int32, len(player))
	for i, score := range player {
		position := getPositionByBinarySearch(uniqueRankeds, score)
		results[i] = int32(position + 1)
	}
	return results
}

// The following code is not part of the solution but is necessary for local functionality.

func printArray(result []int32) {
	for _, resultItem := range result {
		fmt.Printf("%d\n", resultItem)
	}
	fmt.Printf("\n")
}

func main() {
	ranked := []int32{100, 100, 50, 40, 40, 20, 10}
	player := []int32{5, 25, 50, 120}
	result := climbingLeaderboard(ranked, player)
	printArray(result)

	ranked = []int32{100, 90, 90, 80, 75, 60}
	player = []int32{50, 65, 77, 90, 102}
	result = climbingLeaderboard(ranked, player)
	printArray(result)

	return
}
