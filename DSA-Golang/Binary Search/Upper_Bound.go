// Upper Bound
package main

import "fmt"

// UpperBound returns the smallest index i such that nums[i] > x.
// If no such element exists, it returns len(nums).
func upperBound(nums []int, x int) int {
	low := 0
	high := len(nums)
	ans := len(nums)

	for low <= high && low < len(nums) {
		mid := low + (high-low)/2

		if nums[mid] > x {
			ans = mid
			high = mid - 1 // Try to find a smaller index to the left
		} else {
			low = mid + 1 // Move right since nums[mid] <= x
		}
	}

	return ans
}

func main() {
	nums := []int{1, 2, 4, 4, 5, 6, 8}

	fmt.Println(UpperBound(nums, 4))
	fmt.Println(UpperBound(nums, 8))
	fmt.Println(UpperBound(nums, 0))
}
