package main

import (
	"fmt"
)

func arrayOfProducts1(arr []int) []int {

	// Idea.
	// 1. Create filled with 1s three arrays: products, l_products and r_products
	// 2. Initialize variable for accumulating left running product = 1
	// 3. Iterate over array.
	// 3.1 Update the current element of arrays with running product
	// 3.2 re-calculate running product
	// 4. Repeat steps 2 and 3 for r_products array starting from the array end.
	// 5. Multiply l_products and r_products array into products array
	// 6. Return products array
	// 7. Edge cases: empty array, array with one element, array with all ones, array containing Zero(s).

	products := make([]int, len(arr))
	l_products := make([]int, len(arr))
	r_products := make([]int, len(arr))
	for i := range products {
		products[i], l_products[i], r_products[i] = 1, 1, 1
	}

	l_running_product := 1
	for i := 0; i < len(arr); i++ {
		l_products[i] = l_running_product
		l_running_product *= arr[i]
	}

	r_running_product := 1
	for i := len(arr) - 1; i >= 0; i-- {
		r_products[i] = r_running_product
		r_running_product *= arr[i]
	}

	for i := 0; i < len(arr); i++ {
		products[i] = l_products[i] * r_products[i]
	}

	return products
}

func main() {

	arr := []int{5, 1, 4, 2}
	fmt.Println(arrayOfProducts1(arr))
}
