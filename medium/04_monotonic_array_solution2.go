package main

import "fmt"

func monotonicArray2(arr []int) bool {
	// Idea.
	// 1. Having two flags, both initially as True
	// 2. On each iteration two elements. One of flags are
	// definitely become false (unless the whole array is flat)
	// 3. However, if array is not monotonic, both flags will turn as False
	// 4. Finally, return Or of both flags
	is_non_decreasing := true
	is_non_increasing := true

	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			is_non_decreasing = false
		}
		if arr[i] > arr[i-1] {
			is_non_increasing = false
		}
	}
	return is_non_decreasing || is_non_increasing
}

func main() {

	arr := []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}
	fmt.Println(monotonicArray2(arr))
}
