//Single Element in a Sorted Array
//Given a sorted array where every element appears twice except for one element that appears once, we can find it in $O(\log n)$ time and $O(1)$ space using Binary Search.

package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		// Force mid to be even so mid+1 is its potential pair index
		if mid%2 == 1 {
			mid--
		}

		// If nums[mid] == nums[mid+1], the single element is to the right
		if nums[mid] == nums[mid+1] {
			left = mid + 2
		} else {
			// Otherwise, the single element is to the left (including mid)
			right = mid
		}
	}

	return nums[left]
}

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
}
