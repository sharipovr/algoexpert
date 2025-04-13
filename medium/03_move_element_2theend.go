package main

import "fmt"

func moveElement2End(arr []int, target int) []int {
	// 1. create two pointers i and j, at the beginning and end of array
	// 	  which will move towards each other
	// 2. Check if arr at j and i are target numbers
	// 3.1. if a[j] is target move only j
	// 3.2  if a[j] is not and a[i] is -> swap and move j and i
	// 3.3  if a[i] is not and a[i] is not -> move only i

	i := 0
	j := len(arr) - 1
	for i < j {
		if arr[j] == target {
			j--
			continue
		}
		if arr[j] != target {
			if arr[i] == target {
				// swap and move i and j
				arr[i], arr[j] = arr[j], arr[i]
				j--
				i++
			} else {
				i++
			}
		}
	}
	return arr
}

func main() {
	arr := []int{2, 3, 4, 2, 6, 1, 2}
	fmt.Println(moveElement2End(arr, 2))

	arr2 := []int{4, 3, 6, 34, 23, 55, 2, 0, 4}
	fmt.Println(moveElement2End(arr2, 4))
}
