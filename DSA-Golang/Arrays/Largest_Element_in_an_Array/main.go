package main

import "fmt"

func largestElement(nums []int) int {
	largest := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > largest {
			largest = nums[i]
		}
	}

	return largest
}

func main() {
	nums := []int{3, 8, 2, 10, 5}

	result := largestElement(nums)

	fmt.Println("Largest Element:", result)
}
