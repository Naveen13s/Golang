package main

import "fmt"

func rotateArrayByOne(nums []int) {
	n := len(nums)

	if n <= 1 {
		return
	}

	temp := nums[0]

	for i := 1; i < n; i++ {
		nums[i-1] = nums[i]
	}

	nums[n-1] = temp
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	rotateArrayByOne(nums)

	fmt.Println(nums)
}
