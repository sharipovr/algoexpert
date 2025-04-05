package main

import "fmt"

func spiralTraverse(arr [][]int) []int {

	// 1. Declare start and end row and column variables
	// 		as a boundaries of the perimeter of the array
	// 2. Do a loop where sr <= er AND sc <= ec
	// 3. Do four loops by sides of the perimeter,
	//		top->right
	//		right->down
	//		bottom->left	(only if er != sr)
	// 		left->up		(only if sc != ec)
	// 4. update boundaries variables inwards.

	// Edge cases: empty array and flat arrays either single-row
	// or single-column. But this algorithm should handle those
	// without any modification.
	// Any non-squared array is also an edge case and to prevent double-counting
	// we need to add extra conditions for backwards traversals.

	result := []int{}
	sr, er := 0, len(arr)-1
	sc, ec := 0, len(arr[0])-1

	for sr <= er && sc <= ec {
		for col := sc; col <= ec; col++ {
			result = append(result, arr[sr][col])
		}
		for row := sr + 1; row <= er; row++ {
			result = append(result, arr[row][ec])
		}
		if er != sc { // traverse bottom row (if not same as top)
			for col := ec - 1; col >= sc; col-- {
				result = append(result, arr[er][col])
			}
		}

		if sc != ec { // traverse left column (if not same as right)
			for row := er - 1; row > sr; row-- {
				result = append(result, arr[row][sc])
			}
		}
		sr++
		ec--
		er--
		sc++
	}
	return result
}

func main() {

	arr := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 15},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}
	fmt.Println(spiralTraverse(arr))

	arr2 := [][]int{{1, 2, 3}}
	fmt.Println(spiralTraverse(arr2))

	arr3 := [][]int{
		{1, 2, 3, 4},
		{10, 11, 12, 5},
		{9, 8, 7, 6},
	}
	fmt.Println(spiralTraverse(arr3))
}
