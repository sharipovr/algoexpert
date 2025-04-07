package main

import (
	"fmt"
)

func arrayOfProducts2(arr []int) []int {

	// Idea.
	// 1. Create filled with only one arrays for products. Shortener solution.
	// 2. Initialize variable for accumulating left and right running products = 1
	// 3. Iterate over array.
	// 3.1 Update the current element of array with left running product
	// 3.2 re-calculate left running product
	// 4. Go backward over array products starting from the array end.
	//	Update each element by multiplying it on the right running product.
	// 5. Return products array
	// 6. Edge cases: empty array, array with one element, array with all ones, array containing Zero(s).

	products := make([]int, len(arr))

	for i := range products {
		products[i] = 1
	}

	l_running_product := 1
	for i := 0; i < len(arr); i++ {
		products[i] = l_running_product
		l_running_product *= arr[i]
	}

	r_running_product := 1
	for i := len(arr) - 1; i >= 0; i-- {
		products[i] *= r_running_product
		r_running_product *= arr[i]
	}

	return products
}

func main() {

	arr := []int{5, 1, 4, 2}
	fmt.Println(arrayOfProducts2(arr))
}
