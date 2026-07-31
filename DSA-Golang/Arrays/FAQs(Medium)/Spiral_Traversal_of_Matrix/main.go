package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	result := []int{}

	for top <= bottom && left <= right {

		// Left -> Right
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// Top -> Bottom
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// Right -> Left
		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		// Bottom -> Top
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(spiralOrder(matrix))
}

//The Boundary Traversal approach is the preferred interview solution because it visits each element exactly once without requiring an additional visited matrix.
