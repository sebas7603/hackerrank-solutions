package main

import (
	"fmt"
)

// Queen's Attack II
// Problem: https://www.hackerrank.com/challenges/queens-attack-2/problem
// Author:  Sebastián Ayala Suárez
// Since:   29/07/2023

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	res := int32(0)
	left := []int32{0, 0}
	leftDown := []int32{0, 0}
	leftUp := []int32{0, 0}
	right := []int32{0, 0}
	rightDown := []int32{0, 0}
	rightUp := []int32{0, 0}
	down := []int32{0, 0}
	up := []int32{0, 0}

	for _, obstacle := range obstacles {
		if obstacle[1] < c_q {
			// Left
			if obstacle[0] == r_q {
				if left[1] == 0 || obstacle[1] > left[1] {
					left = obstacle
					continue
				}
			}

			// Diag Left-Down
			if obstacle[0] == r_q-(c_q-obstacle[1]) {
				if leftDown[1] == 0 || obstacle[1] > leftDown[1] {
					leftDown = obstacle
					continue
				}
			}

			// Diag Left-Up
			if obstacle[0] == r_q+(c_q-obstacle[1]) {
				if leftUp[1] == 0 || obstacle[1] > leftUp[1] {
					leftUp = obstacle
					continue
				}
			}
		} else if obstacle[1] > c_q {
			// Right
			if obstacle[0] == r_q {
				if right[1] == 0 || obstacle[1] < right[1] {
					right = obstacle
					continue
				}
			}

			// Diag Right-Down
			if obstacle[0] == r_q-(obstacle[1]-c_q) {
				if rightDown[1] == 0 || obstacle[1] < rightDown[1] {
					rightDown = obstacle
					continue
				}
			}

			// Diag Right-Up
			if obstacle[0] == r_q+(obstacle[1]-c_q) {
				if rightUp[1] == 0 || obstacle[1] < rightUp[1] {
					rightUp = obstacle
					continue
				}
			}
		} else {
			// Down
			if obstacle[0] < r_q {
				if down[0] == 0 || obstacle[0] > down[0] {
					down = obstacle
					continue
				}
			}

			// Up
			if obstacle[0] > r_q {
				if up[0] == 0 || obstacle[0] < up[0] {
					up = obstacle
					continue
				}
			}
		}
	}

	if left[1] == 0 {
		res += c_q - 1
	} else {
		res += c_q - left[1] - 1
	}

	if leftDown[1] == 0 {
		if c_q < r_q {
			res += c_q - 1
		} else {
			res += r_q - 1
		}
	} else {
		res += c_q - leftDown[1] - 1
	}

	if leftUp[1] == 0 {
		if c_q-1 < n-r_q {
			res += c_q - 1
		} else {
			res += n - r_q
		}
	} else {
		res += c_q - leftUp[1] - 1
	}

	if right[1] == 0 {
		res += n - c_q
	} else {
		res += right[1] - c_q - 1
	}

	if rightDown[1] == 0 {
		if n-c_q < r_q-1 {
			res += n - c_q
		} else {
			res += r_q - 1
		}
	} else {
		res += rightDown[1] - c_q - 1
	}

	if rightUp[1] == 0 {
		if c_q > r_q {
			res += n - c_q
		} else {
			res += n - r_q
		}
	} else {
		res += rightUp[1] - c_q - 1
	}

	if down[0] == 0 {
		res += r_q - 1
	} else {
		res += r_q - down[0] - 1
	}

	if up[0] == 0 {
		res += n - r_q
	} else {
		res += up[0] - r_q - 1
	}

	return res
}

// The following code is not part of the solution but is necessary for local functionality.

func main() {
	obstacles := [][]int32{}
	res := queensAttack(4, 0, 4, 4, obstacles)
	fmt.Printf("%d\n", res)

	obstacles2 := [][]int32{
		{5, 5},
		{4, 2},
		{2, 3},
	}
	res2 := queensAttack(5, 3, 4, 3, obstacles2)
	fmt.Printf("%d\n", res2)

	res3 := queensAttack(1, 0, 1, 1, obstacles)
	fmt.Printf("%d\n", res3)

	return
}
