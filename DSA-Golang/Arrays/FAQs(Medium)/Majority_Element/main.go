// Moore's Voting Algorithm is the expected optimal solution in interviews because it achieves both linear time and constant extra space.
package main

import "fmt"

func majorityElement(nums []int) int {
	candidate := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}

	fmt.Println("Majority Element:", majorityElement(nums))
}
