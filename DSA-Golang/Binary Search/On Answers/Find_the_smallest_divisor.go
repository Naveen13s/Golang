//Find the smallest divisor
/* Given an array of integers nums and an integer limit as the threshold value,
   find the smallest positive integer divisor such that upon dividing all the elements of the array by this divisor,
   the sum of the division results is less than or equal to the threshold value.
*/

package main

import "fmt"

func smallestDivisor(nums []int, limit int) int {
	left, right := 1, 0

	// Find the maximum element in nums to set the upper bound
	for _, num := range nums {
		if num > right {
			right = num
		}
	}

	ans := right

	for left <= right {
		mid := left + (right-left)/2

		if calculateSum(nums, mid) <= limit {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

// Helper function to calculate total sum after ceiling division
func calculateSum(nums []int, divisor int) int {
	sum := 0
	for _, num := range nums {
		// Equivalent to ceil(num / divisor) using integer math
		sum += (num + divisor - 1) / divisor
	}
	return sum
}

func main() {
	fmt.Println(smallestDivisor([]int{1, 2, 5, 9}, 6))
	fmt.Println(smallestDivisor([]int{44, 22, 33, 11, 1}, 5))
}
