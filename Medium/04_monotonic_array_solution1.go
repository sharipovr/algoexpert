package main

import "fmt"

func monotonicArray1(arr []int) bool {
	// Idea:
	// 1. We need a flag (var direction (1, 0, -1) showing direction is determined
	// 2. If direction = 0, we try to find out the direction on every iteration.
	// 3. If direction <> 0, we compare current element and previous one:
	// 		if comparison satisfies the direction, or it's neutral we proceed,
	//		otherwise we return false
	// 4. If we exited loop we return true
	// 5. Edge cases: if len of array is less than 3.

	if len(arr) < 3 {
		return true // Edge case
	}

	direction := 0 // 1, -1, 0
	for i := 1; i < len(arr); i++ {
		if direction == 0 {
			if arr[i] < arr[i-1] {
				direction = -1
			} else if arr[i] > arr[i-1] {
				direction = 1
			}
		}

		if arr[i] < arr[i-1] && direction == 1 {
			return false
		}
		if arr[i] > arr[i-1] && direction == -1 {
			return false
		}
	}

	return true
}

func main() {

	arr := []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}
	fmt.Println(monotonicArray1(arr))

	arr = []int{-1, 2}
	fmt.Println(monotonicArray1(arr))
}
