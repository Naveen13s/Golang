package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	maxCount := 0

	for _, num := range nums {
		if num == 1 {
			count++

			if count > maxCount {
				maxCount = count
			}
		} else {
			count = 0
		}
	}

	return maxCount
}

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}

	result := findMaxConsecutiveOnes(nums)

	fmt.Println("Maximum Consecutive 1s:", result)
}
