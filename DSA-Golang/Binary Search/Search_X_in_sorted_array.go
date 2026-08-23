// Search X in sorted array
package main

import "fmt"

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2 // Prevents potential integer overflow

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1 // Target lies in the right half
		} else {
			high = mid - 1 // Target lies in the left half
		}
	}

	return -1 // Target not found
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}

	fmt.Println(search(nums, 9)) // Output: 4
	fmt.Println(search(nums, 2)) // Output: -1
}
