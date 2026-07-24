package main

import "fmt"

func linearSearch(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}

	return -1
}

func main() {
	nums := []int{2, 3, 4, 3, 5}
	target := 3

	result := linearSearch(nums, target)

	fmt.Println("Smallest Index:", result)
}
