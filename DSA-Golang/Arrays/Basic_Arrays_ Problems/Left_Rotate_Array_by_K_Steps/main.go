package main

import "fmt"

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func rotateLeft(nums []int, k int) {
	n := len(nums)

	if n == 0 {
		return
	}

	k = k % n

	if k == 0 {
		return
	}

	// Reverse first k elements
	reverse(nums, 0, k-1)

	// Reverse remaining elements
	reverse(nums, k, n-1)

	// Reverse entire array
	reverse(nums, 0, n-1)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	rotateLeft(nums, k)

	fmt.Println(nums)
}
