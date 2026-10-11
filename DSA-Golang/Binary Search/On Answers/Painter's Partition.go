//Painter's Partition

package main

import "fmt"

func paint(A int, B int, C []int) int {
	var left int64 = 0
	var right int64 = 0

	for _, length := range C {
		if int64(length) > left {
			left = int64(length)
		}
		right += int64(length)
	}

	ans := right

	// Binary search for the minimum possible max length a painter handles
	for left <= right {
		mid := left + (right-left)/2
		if canPaint(C, A, mid) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// Calculate total time and apply modulo 10000003
	mod := int64(10000003)
	totalTime := (ans % mod) * (int64(B) % mod) % mod
	return int(totalTime)
}

// Helper function to check if boards can be painted within 'A' painters given a max length limit
func canPaint(C []int, A int, maxAllowedLength int64) bool {
	painters := 1
	var currentLength int64 = 0

	for _, length := range C {
		if int64(length) > maxAllowedLength {
			return false
		}
		if currentLength+int64(length) > maxAllowedLength {
			painters++
			currentLength = int64(length)
		} else {
			currentLength += int64(length)
		}
	}

	return painters <= A
}

func main() {
	A := 2
	B := 5
	C := []int{1, 10}

	result := paint(A, B, C)
	fmt.Printf("Minimum time required: %d\n", result)
}
