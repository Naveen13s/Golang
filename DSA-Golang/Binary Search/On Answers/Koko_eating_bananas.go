//Koko eating bananas

package main

import "fmt"

func minEatingSpeed(nums []int, h int) int {
	left, right := 1, 0

	// Find the maximum pile size to set the upper bound
	for _, pile := range nums {
		if pile > right {
			right = pile
		}
	}

	ans := right

	for left <= right {
		mid := left + (right-left)/2

		// Calculate total hours needed at speed 'mid' inline
		totalHours := 0
		for _, pile := range nums {
			totalHours += (pile + mid - 1) / mid
		}

		// Check if we can finish within h hours
		if totalHours <= h {
			ans = mid // Valid speed, try finding a smaller one
			right = mid - 1
		} else {
			left = mid + 1 // Too slow, need a higher speed
		}
	}

	return ans
}

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
}
