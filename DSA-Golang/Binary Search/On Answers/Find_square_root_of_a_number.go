//Find square root of a number
//Given a positive integer n. Find and return its square root. If n is not a perfect square, then return the floor value of sqrt(n).

package main

import "fmt"

func mySqrt(n int) int {
	if n == 0 {
		return 0
	}

	left, right := 1, n
	ans := 1

	for left <= right {
		mid := left + (right-left)/2

		// Prevent overflow: equivalent to mid * mid <= n
		if mid <= n/mid {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8)) // Output: 2 (floor of 2.828...)
	fmt.Println(mySqrt(16))
}
