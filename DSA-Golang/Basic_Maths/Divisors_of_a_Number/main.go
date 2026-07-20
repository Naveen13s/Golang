package main

import (
	"fmt"
	"sort"
)

func divisors(n int) []int {
	var result []int

	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			result = append(result, i)

			if i != n/i {
				result = append(result, n/i)
			}
		}
	}

	sort.Ints(result)

	return result
}

func main() {
	n := 12

	fmt.Println(divisors(n))
}
