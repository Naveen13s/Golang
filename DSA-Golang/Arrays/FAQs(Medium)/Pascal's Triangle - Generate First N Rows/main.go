package main

import "fmt"

func generatePascalTriangle(n int) [][]int {
	triangle := make([][]int, n)

	for i := 0; i < n; i++ {
		triangle[i] = make([]int, i+1)

		// First and last element
		triangle[i][0] = 1
		triangle[i][i] = 1

		// Fill middle elements
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}

func main() {
	n := 5

	result := generatePascalTriangle(n)

	for _, row := range result {
		fmt.Println(row)
	}
}
