package main

import "fmt"

func pascalRow(r int) []int {
	n := r - 1

	row := make([]int, 0, r)

	current := 1
	row = append(row, current)

	for i := 0; i < n; i++ {
		current = current * (n - i) / (i + 1)
		row = append(row, current)
	}

	return row
}

func main() {
	r := 5

	fmt.Println(pascalRow(r))
}

//Alternative Approach (Generate Entire Triangle)
//Another approach is to build Pascal's Triangle row by row and return the last row.

func pascalRow(r int) []int {
	triangle := make([][]int, r)

	for i := 0; i < r; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle[r-1]
}

//For printing a single row, avoid generating the entire triangle. The iterative combination method is the expected optimal solution because it runs in O(r) time.
