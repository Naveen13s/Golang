//Search in rotated sorted array-I
/* Given an integer array nums, sorted in ascending order (with distinct values) and a target value k. The array is rotated at some pivot point that is unknown.
   Find the index at which k is present and if k is not present return -1.
*/

package main

import "fmt"

func search(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == k {
			return mid
		}
		// Check if the left half is sorted
		if nums[left] <= nums[mid] {
			// Check if k lies within the sorted left range
			if nums[left] <= k && k < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// Otherwise, the right half must be sorted
			// Check if k lies within the sorted right range
			if nums[mid] < k && k <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}

	fmt.Println(search(nums, 0))
	fmt.Println(search(nums, 3))
}
