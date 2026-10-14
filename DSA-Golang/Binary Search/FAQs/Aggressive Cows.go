//Aggressive Cows
/* Given an array nums of size n, which denotes the positions of stalls, and an integer k, which denotes the number of aggressive cows, assign stalls to k cows such that the minimum distance between any two cows is the maximum possible. Find the maximum possible minimum distance.*/

package main

import (
	"fmt"
	"sort"
)

func maxDistance(nums []int, k int) int {
	sort.Ints(nums)
	left := 1
	right := nums[len(nums)-1] - nums[0]
	ans := 0
	for left <= right {
		mid := left + (right-left)/2

		if canPlaceCows(nums, k, mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

// Helper function to check if we can place 'k' cows with at least 'minDist' apart
func canPlaceCows(nums []int, k int, minDist int) bool {
	cowsPlaced := 1
	lastPos := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]-lastPos >= minDist {
			cowsPlaced++
			lastPos = nums[i]
			if cowsPlaced >= k {
				return true
			}
		}
	}

	return false
}

func main() {
	nums := []int{1, 2, 4, 8, 9}
	k := 3

	result := maxDistance(nums, k)
	fmt.Printf("Maximum possible minimum distance: %d\n", result)
}
