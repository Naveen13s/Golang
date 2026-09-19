//Find minimum in Rotated Sorted Array
//Given an integer array nums of size N, sorted in ascending order with distinct values, and then rotated an unknown number of times (between 1 and N), find the minimum element in the array.

package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		// agr mid element is greater than the rightmost element,the minimum MUST be in the right half (mid + 1 to right)
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			// nahi to, the minimum is in the left half, including mid
			right = mid
		}
	}

	return nums[left]
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2})) 
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2})) 
	fmt.Println(findMin([]int{11, 13, 15, 17})) 