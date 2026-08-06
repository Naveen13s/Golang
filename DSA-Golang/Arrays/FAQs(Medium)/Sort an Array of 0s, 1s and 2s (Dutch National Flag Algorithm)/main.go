package main

import "fmt"

func sortColors(nums []int) {
	low, mid := 0, 0
	high := len(nums) - 1

	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++

		case 1:
			mid++

		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}

	sortColors(nums)

	fmt.Println(nums)
}

//The Dutch National Flag Algorithm, proposed by Edsger Dijkstra, is the expected interview solution because it sorts the array in a single pass, in-place, and uses constant extra space.
