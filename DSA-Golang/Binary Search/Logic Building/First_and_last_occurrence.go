//First and last occurrence
//Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value. If the target is not found in the array, return [-1, -1].

package main

import "fmt"

func searchRange(nums []int, target int) []int {
	return []int{
		findBound(nums, target, true),  // Find first position
		findBound(nums, target, false), // Find last position
	}
}

// Helper function to find either the first or last index of the target
func findBound(nums []int, target int, isFirst bool) int {
	left, right := 0, len(nums)-1
	bound := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			bound = mid
			if isFirst {
				right = mid - 1 // Keep searching left for the first occurrence
			} else {
				left = mid + 1 // Keep searching right for the last occurrence
			}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return bound
}
func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums, 8))
	fmt.Println(searchRange(nums, 6))
	fmt.Println(searchRange([]int{}, 0))
}
