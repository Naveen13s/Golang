// Lower Bound
package main

import "fmt"

func lowerBound(nums []int, x int) int {
	low, high := 0, len(nums)-1
	ans := len(nums) // Default answer if no element is >= x

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] >= x {
			ans = mid      // Potential lower bound found
			high = mid - 1 // Search left to find a smaller valid index
		} else {
			low = mid + 1 // Value too small, search right
		}
	}

	return ans
}

func main() {
	nums := []int{1, 2, 4, 4, 5, 6, 8, 10}

	fmt.Println(lowerBound(nums, 4))  // Output: 2 (first index where val >= 4)
	fmt.Println(lowerBound(nums, 7))  // Output: 6 (index of 8, the first val >= 7)
	fmt.Println(lowerBound(nums, 11)) // Output: 8 (len(nums), since 11 > all elements)
}
