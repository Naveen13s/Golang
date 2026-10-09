//Capacity to Ship Packages Within D Days
/* You are given an array weights where weights[i] represents the weight of the i-th package on a conveyor belt. All the packages must be shipped in the order given from one port to another within days days.
   Each day, the ship can carry a contiguous sequence of packages, as long as the total weight does not exceed its maximum capacity.
   Your task is to find the minimum possible capacity of the ship so that all packages can be shipped within the given number of days.
*/

package main

import "fmt"

func shipWithinDays(weights []int, days int) int {
	left, right := 0, 0
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}

	ans := right

	// Binary search for the minimum capacity
	for left <= right {
		mid := left + (right-left)/2
		if canShip(weights, days, mid) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

// Helper function to check if packages can be shipped within 'days' with a given 'capacity'
func canShip(weights []int, days int, capacity int) bool {
	currentDays := 1
	currentWeight := 0

	for _, w := range weights {
		if currentWeight+w > capacity {
			currentDays++
			currentWeight = w
		} else {
			currentWeight += w
		}
	}

	return currentDays <= days
}

func main() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	days := 5

	result := shipWithinDays(weights, days)
	fmt.Printf("Minimum capacity needed: %d\n", result)
}
