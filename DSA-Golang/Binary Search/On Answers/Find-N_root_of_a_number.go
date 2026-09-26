//Find Nth root of a number
//Given two numbers N and M, find the Nth root of M. The Nth root of a number M is defined as a number X such that when X is raised to the power of N, it equals M. If the Nth root is not an integer, return -1.

package main

import "fmt"

func nthRoot(n int, m int) int {
	left, right := 1, m

	for left <= right {
		mid := left + (right-left)/2
		val := power(mid, n, m)

		if val == 1 {
			return mid
		} else if val == 2 {
			right = mid - 1
		} else {
			left = mid + 1 // mid^n is too small, search right
		}
	}

	return -1 // No integer Nth root exists
}

// Helper function to safely calculate base^exp and compare with target (m)
// Returns: 0 if base^exp < target, 1 if equal, 2 if base^exp > target
func power(base, exp, target int) int {
	res := 1
	for i := 0; i < exp; i++ {
		// Prevent overflow: check if res * base exceeds target
		if target/base < res {
			return 2
		}
		res *= base
	}

	if res == target {
		return 1
	}
	return 0
}

func main() {
	fmt.Println(nthRoot(3, 27))
	fmt.Println(nthRoot(4, 69))
	fmt.Println(nthRoot(2, 9))
}
