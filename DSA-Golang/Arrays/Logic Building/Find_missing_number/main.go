package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)

	expectedSum := n * (n + 1) / 2

	actualSum := 0

	for _, num := range nums {
		actualSum += num
	}

	return expectedSum - actualSum
}

func main() {
	nums := []int{3, 0, 1}

	fmt.Println("Missing Number:", missingNumber(nums))
}



// Another Approach- Bit Manipulation - XOR

package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)

	xor := n

	for i := 0; i < n; i++ {
		xor ^= i
		xor ^= nums[i]
	}

	return xor
}

func main() {
	nums := []int{3, 0, 1}

	fmt.Println("Missing Number:", missingNumber(nums))
}