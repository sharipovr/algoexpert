package main

import "fmt"

func longestPeak(arr []int) int {
	// The idea:
	// 1. Iterate over array
	// 2. Looking the tip of the peak.
	// 3. If it's not peak, iterate over, otherwise:
	//   look for how "long" the peak might be by looking to the left and
	// 	 looking to the right.
	//   We are doing two loops (kind of while) until indexes are in limits of array
	//	 and slopes are going down. Eventually we have left and right indexes
	// 4. We can calculate length of peak by subtracting indexes and -1
	// 5. Finally, if this length is greater the longest_length we update longest_length
	// 6. We update i to be equal to the right index of exploring right slope "procedure"
	// 7. After exit from loop we return longest_length
	// 8. Edge cases:
	// 		if there are less than 3 elements in the array
	//		if array is monotonic or constant

	longest_peak_len := 0

	for i := 1; i < len(arr)-1; {
		var is_peak bool
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			is_peak = true
		}

		if !is_peak {
			i++
			continue
		}

		// Looking to the sides

		left_idx := i - 2
		for left_idx >= 0 && arr[left_idx] < arr[left_idx+1] {
			left_idx--
		}
		right_idx := i + 2
		for right_idx < len(arr) && arr[right_idx] < arr[right_idx-1] {
			right_idx++
		}

		current_peak_len := right_idx - left_idx - 1

		if current_peak_len > longest_peak_len {
			longest_peak_len = current_peak_len
		}
		i = right_idx
	}

	return longest_peak_len
}

func main() {
	arr := []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}
	fmt.Println(longestPeak(arr))
}
