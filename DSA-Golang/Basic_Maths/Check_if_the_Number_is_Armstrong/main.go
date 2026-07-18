//Best Approach own integer power function

import (
	"fmt"
	"math"
)
func power(base, exp int) int {
	result := 1

	for i := 0; i < exp; i++ {
		result *= base
	}

	return result
}

func isArmstrong(n int) bool {
	if n == 0 {
		return true
	}

	original := n
	count := 0
	temp := n

	// Count digits
	for temp > 0 {
		count++
		temp /= 10
	}

	sum := 0
	temp = n

	// Calculate sum
	for temp > 0 {
		digit := temp % 10
		sum += power(digit, count)
		temp /= 10
	}

	return sum == original
}

// 2nd  Approach
package main

import (
	"fmt"
	"math"
)

func isArmstrong(n int) bool {
	if n == 0 {
		return true
	}

	original := n
	count := 0
	temp := n

	// Count digits
	for temp > 0 {
		count++
		temp /= 10
	}

	sum := 0
	temp = n

	// Calculate sum of digits raised to the power of count
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(count)))
		temp /= 10
	}

	return sum == original
}

func main() {
	n := 153

	fmt.Println(isArmstrong(n))
}