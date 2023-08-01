package main

import (
	"fmt"
)

// Fraudulent Activity Notifications
// Problem: https://www.hackerrank.com/challenges/fraudulent-activity-notifications/problem
// Author:  Sebastián Ayala Suárez
// Since:   31/07/2023
// Note:    getPositionByBinarySearch uses binary search to find the position of a score

func getPositionByBinarySearch(array []int32, value int32) int {
	first, last := 0, len(array)-1
	for first <= last {
		mid := (first + last) / 2
		if array[mid] == value {
			return mid
		} else if array[mid] < value {
			last = mid - 1
		} else {
			first = mid + 1
		}
	}
	return first
}

// Insert a value into an array at a given index, preserving the order of existing elements
func insert(array []int32, index int32, value int32) []int32 {
	if int32(len(array)) == index {
		return append(array, value)
	}
	array = append(array[:index+1], array[index:]...)
	array[index] = value
	return array
}

// Remove a value into an array at a given index, preserving the order of existing elements
func remove(array []int32, index int32) []int32 {
	return append(array[:index], array[index+1:]...)
}

func activityNotifications(expenditure []int32, d int32) int32 {
	notifications := int32(0)
	trailingExpenitures := []int32{expenditure[0]}

	// Initial sort
	for i := int32(1); i < d; i++ {
		done := false
		j := int32(0)
		for !done && j < i {
			if expenditure[i] >= trailingExpenitures[j] {
				done = true
				trailingExpenitures = insert(trailingExpenitures, j, expenditure[i])
			}
			j++
		}
		if !done {
			trailingExpenitures = insert(trailingExpenitures, j, expenditure[i])
		}
	}

	// Start eval notifications
	var median float64
	for i := d; i < int32(len(expenditure)); i++ {
		median = float64(trailingExpenitures[d/2])
		if d%2 == 0 {
			median = (median + float64(trailingExpenitures[(d/2)-1])) / 2
		}

		if expenditure[i] >= int32(median*2) {
			notifications++
		}

		if expenditure[i-d] != expenditure[i] && i != int32(len(expenditure)-1) {
			positionToRemove := getPositionByBinarySearch(trailingExpenitures, expenditure[i-d])
			trailingExpenitures = remove(trailingExpenitures, int32(positionToRemove))

			positionToAdd := getPositionByBinarySearch(trailingExpenitures, expenditure[i])
			trailingExpenitures = insert(trailingExpenitures, int32(positionToAdd), expenditure[i])
		}
	}

	return notifications
}

// The following code is not part of the solution but is necessary for local functionality.

func main() {
	expenditure1 := []int32{2, 3, 4, 2, 3, 6, 8, 4, 5}
	res1 := activityNotifications(expenditure1, 5)
	fmt.Printf("%d\n", res1)

	expenditure2 := []int32{1, 2, 3, 4, 4}
	res2 := activityNotifications(expenditure2, 4)
	fmt.Printf("%d\n", res2)

	expenditure3 := []int32{10, 20, 30, 40, 50}
	res3 := activityNotifications(expenditure3, 3)
	fmt.Printf("%d\n", res3)
	return
}
