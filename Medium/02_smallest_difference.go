package main

import (
	"fmt"
	"sort"
)

func smallestDifference1(arr1, arr2 []int) []int {
	// 1. Sort both arrays in place
	// {-1, 3, 5, 10, 20, 28}
	// {15, 17, 26, 134, 135}
	// 2. Start wiht first element in both array, create i and j = 0
	// 3. Compare arr1[i] and arr2[j] (a1 and a2 as a separate preserved variables)
	// 3.1 If they match we are done, return this pair
	// 3.2 If a1 < a2 -> we have diff and moving i will potentially decrease it, cause arr1 is sorted
	// 		if diff less than smalles_diff save diff as smallest_diff and the pair
	// 3.3 If a2 < a1 -> we have a diff and now moving j will probably decrease it, because arr2 is sorted
	// 		save diff as smallest_diff, and pair

	sort.Ints(arr1)
	sort.Ints(arr2)
	i := 0
	j := 0
	smallest_diff := 2_000_000_000
	curr_diff := 0
	result := make([]int, 2)

	for i < len(arr1) && j < len(arr2) {
		a1 := arr1[i]
		a2 := arr2[j]
		// Happy path, match
		if a1 == a2 {
			return []int{a1, a2}
		}

		if a1 < a2 {
			curr_diff = a2 - a1
			i++
		} else {
			curr_diff = a1 - a2
			j++
		}

		if curr_diff < smallest_diff {
			smallest_diff = curr_diff
			result = []int{a1, a2}
		}

	}
	return result
}

func main() {
	arrayOne := []int{-1, 5, 10, 20, 28, 3}
	arrayTwo := []int{26, 134, 135, 15, 17}
	result := smallestDifference1(arrayOne, arrayTwo)
	fmt.Println(result) // Output should be: [28, 26]
}
