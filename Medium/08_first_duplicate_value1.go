package main

import "fmt"

func firstDuplicateValue1(arr []int) int {

	// Idea: naive approach with two loop approach
	// Init minIndex variable = n, which is len(arr)
	// In inner loop we check if a[i] == a[j]
	// if so, check if j < minIndex and update it
	// return arr[minIndex] unless minIndex has not changed
	// otherwise return -1
	// Edge case less than two elements in array
	// Complexity: time O(N^2), time O(1)
	n := len(arr)
	if n < 2 {
		return -1
	}
	minimumIndex := n
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] == arr[j] && j < minimumIndex {
				minimumIndex = j
			}
		}
	}
	if minimumIndex != n {
		return arr[minimumIndex]
	}
	return -1
}

func main() {

	arr := []int{2, 1, 5, 2, 3, 3, 4}
	fmt.Println(firstDuplicateValue1(arr))
}
