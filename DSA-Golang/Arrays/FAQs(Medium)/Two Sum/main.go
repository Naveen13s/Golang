package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, found := hashMap[complement]; found {
			return []int{index, i}
		}

		hashMap[num] = i
	}

	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))
}

//The Hash Map approach is the expected interview solution because it finds the answer in a single pass while preserving the original indices.
