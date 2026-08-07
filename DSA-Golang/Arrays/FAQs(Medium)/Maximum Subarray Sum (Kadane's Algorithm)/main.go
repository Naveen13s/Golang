package main

import "fmt"

func maxSubArray(nums []int) int {
	currentSum := 0
	maxSum := nums[0]

	for _, num := range nums {
		currentSum += num

		if currentSum > maxSum {
			maxSum = currentSum
		}

		if currentSum < 0 {
			currentSum = 0
		}
	}

	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println("Maximum Subarray Sum:", maxSubArray(nums))
}
