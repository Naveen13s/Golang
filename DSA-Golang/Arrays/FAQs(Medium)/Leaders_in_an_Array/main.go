package main

import "fmt"

func leaders(nums []int) []int {
	n := len(nums)

	if n == 0 {
		return []int{}
	}

	result := []int{}
	maxRight := nums[n-1]

	result = append(result, maxRight)

	for i := n - 2; i >= 0; i-- {
		if nums[i] > maxRight {
			result = append(result, nums[i])
			maxRight = nums[i]
		}
	}

	// Reverse the result
	for left, right := 0, len(result)-1; left < right; {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}

	return result
}

func main() {
	nums := []int{16, 17, 4, 3, 5, 2}

	fmt.Println(leaders(nums))
}
