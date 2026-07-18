package main

import "fmt"

func factorial(n int) int {
	result := 1

	for i := 2; i <= n; i++ {
		result *= i
	}

	return result
}

func main() {
	n := 5

	fmt.Println("Factorial:", factorial(n))
}
