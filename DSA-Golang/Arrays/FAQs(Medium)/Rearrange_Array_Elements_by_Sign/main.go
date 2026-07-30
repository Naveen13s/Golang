package main

import "fmt"

func rearrangeArray(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	pos := 0
	neg := 1

	for _, num := range nums {
		if num > 0 {
			result[pos] = num
			pos += 2
		} else {
			result[neg] = num
			neg += 2
		}
	}

	return result
}

func main() {
	nums := []int{3, 1, -2, -5, 2, -4}

	fmt.Println(rearrangeArray(nums))
}
