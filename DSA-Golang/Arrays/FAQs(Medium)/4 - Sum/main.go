package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	result := [][]int{}
	n := len(nums)

	for i := 0; i < n-3; i++ {

		// Skip duplicate values for i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n-2; j++ {

			// Skip duplicate values for j
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left := j + 1
			right := n - 1

			for left < right {
				sum := int64(nums[i]) +
					int64(nums[j]) +
					int64(nums[left]) +
					int64(nums[right])

				if sum == int64(target) {
					result = append(result, []int{
						nums[i],
						nums[j],
						nums[left],
						nums[right],
					})

					// Skip duplicates
					for left < right && nums[left] == nums[left+1] {
						left++
					}

					for left < right && nums[right] == nums[right-1] {
						right--
					}

					left++
					right--

				} else if sum < int64(target) {
					left++
				} else {
					right--
				}
			}
		}
	}

	return result
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0

	result := fourSum(nums, target)

	fmt.Println(result)
}

//A useful pattern to remember is: Two Sum → Hash Map, 3Sum → fix 1 + two pointers, 4Sum → fix 2 + two pointers.
