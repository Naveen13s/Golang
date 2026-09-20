//Find out how many times the array is rotated
/*Given an integer array nums of size n, sorted in ascending order with distinct values. The array has been right rotated an unknown number of times, between 0 and n-1 (including).
Determine the number of rotations performed on the array.
*/
package main

import "fmt"

func findKRotation(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		// Minimum element is in the right half
		if nums[mid] > nums[right] {
			left = mid + 1
		} else { // Minimum element is in the left half (including mid)
			right = mid
		}
	}
	// 'left' is the index of the minimum element
	return left
}

func main() {
	fmt.Println(findKRotation([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findKRotation([]int{3, 4, 5, 1, 2}))
	fmt.Println(findKRotation([]int{1, 2, 3, 4, 5}))
}
