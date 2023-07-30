package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Time Conversion
// Problem: https://www.hackerrank.com/challenges/time-conversion/problem
// Author:  Sebastián Ayala Suárez
// Since:   29/07/2023

func timeConversion(s string) string {
	if strings.Contains(s, "AM") {
		hourAm := s[0:2]
		newTimeAm := strings.TrimSuffix(s, "AM")
		if hourAm == "12" {
			return fmt.Sprintf("00%s", strings.TrimPrefix(newTimeAm, hourAm))
		}
		return newTimeAm
	}

	hour := s[0:2]
	newTime := strings.TrimSuffix(s, "PM")
	if hour == "12" {
		return newTime
	}
	newHour, _ := strconv.Atoi(hour)
	return fmt.Sprintf("%d%s", newHour+12, strings.TrimPrefix(newTime, hour))
}

// The following code is not part of the solution but is necessary for local functionality.

func main() {
	res := timeConversion("07:05:45PM")
	fmt.Println(res)

	res = timeConversion("12:01:00PM")
	fmt.Println(res)

	res = timeConversion("12:01:00AM")
	fmt.Println(res)

	return
}
