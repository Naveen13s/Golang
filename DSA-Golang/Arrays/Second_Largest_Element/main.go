package main

import "fmt"

func secondLargestElement(nums []int) int {
	if len(nums) < 2 {
		return -1
	}

	largest := nums[0]
	secondLargest := -1

	for i := 1; i < len(nums); i++ {
		if nums[i] > largest {
			secondLargest = largest
			largest = nums[i]
		} else if nums[i] < largest && nums[i] > secondLargest {
			secondLargest = nums[i]
		}
	}

	return secondLargest
}

func main() {
	nums := []int{8, 8, 7, 6, 5}

	fmt.Println("Second Largest Element:", secondLargestElement(nums))
}
