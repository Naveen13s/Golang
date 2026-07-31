package main

import "fmt"

func pascalElement(r, c int) int {
	n := r - 1
	k := c - 1

	if k > n-k {
		k = n - k
	}

	result := 1

	for i := 0; i < k; i++ {
		result = result * (n - i)
		result = result / (i + 1)
	}

	return result
}

func main() {
	r := 5
	c := 3

	fmt.Println(pascalElement(r, c))
}
