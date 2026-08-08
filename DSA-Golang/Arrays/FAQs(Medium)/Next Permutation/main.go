package main

import "fmt"

func nextPermutation(nums []int) {
	n := len(nums)

	// Step 1: Find pivot
	pivot := -1
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			pivot = i
			break
		}
	}

	// Step 2: If pivot exists, swap with successor
	if pivot != -1 {
		for i := n - 1; i > pivot; i-- {
			if nums[i] > nums[pivot] {
				nums[i], nums[pivot] = nums[pivot], nums[i]
				break
			}
		}
	}

	// Step 3: Reverse suffix
	left := pivot + 1
	right := n - 1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func main() {
	nums := []int{1, 2, 3}

	nextPermutation(nums)

	fmt.Println(nums)
}

/*
This is one of the most frequently asked array interview questions because it combines greedy thinking with array manipulation.

Remember this sequence:
Find Pivot → Find Successor → Swap → Reverse

This pattern guarantees the next lexicographically greater permutation in O(N) time and O(1) extra space.
*/
