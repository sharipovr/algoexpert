package main

import (
	"fmt"
	"sort"
)

func threenums_sum(arr []int, target int) [][]int {
	result := [][]int{}
	// We want to return the triple from array i.e. [5, -3, -2] if target = 0, cause this sums up to 0.
	// Idea:
	// Sort an array
	// 1. Do first loop with index 'i' from the left to right. Not necessary to the last element, but to n - 2
	// 2. Use two pointers: (left) l = i+1, and (right) r = len(arr)-1 and move them towards each other
	// 3. Calculate sum of triplet and if it equals to target save to result array
	//    if sum is less than target, we have not enough values, so move l
	//	  if sum is greater than target we have too much values, so move r
	// 4. Return result
	// 5. Edge cases:
	// 	  Length of array is less 3

	if len(arr) < 3 {
		fmt.Println("Array must: have at least three elements!")
		return result
	}

	sort.Ints(arr)
	for i := 0; i < len(arr)-2; i++ {
		l := i + 1
		r := len(arr) - 1

		for l < r {
			candidate := arr[i] + arr[l] + arr[r]
			if candidate == target {
				result = append(result, []int{arr[i], arr[l], arr[r]})
			}
			if candidate < target {
				l++
			} else {
				r--
			}
		}
	}
	return result
}

func main() {

	arr := []int{1, 2, -2, 6, 5, 3, 8, -3} // {-3, -2, 1, 2, 3, 5, 6, 8}
	target := 0

	arr1 := []int{1, 2} // Edge case

	arr2 := []int{1, 2, 3} // No match
	fmt.Println(threenums_sum(arr, target))
	fmt.Println(threenums_sum(arr1, target))
	fmt.Println(threenums_sum(arr2, 7))
}
