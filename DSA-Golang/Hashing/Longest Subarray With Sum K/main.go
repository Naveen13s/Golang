package main

import "fmt"

func longestSubarray(nums []int, k int) int {
	prefixSum := 0
	maxLength := 0

	// prefixSum -> first index
	prefixMap := make(map[int]int)

	// Important for subarrays starting from index 0
	prefixMap[0] = -1

	for i, num := range nums {
		prefixSum += num

		// Check if required prefix sum exists
		if index, found := prefixMap[prefixSum-k]; found {
			length := i - index

			if length > maxLength {
				maxLength = length
			}
		}

		// Store only the first occurrence
		if _, found := prefixMap[prefixSum]; !found {
			prefixMap[prefixSum] = i
		}
	}

	return maxLength
}

func main() {
	nums := []int{10, 5, 2, 7, 1, 9}
	k := 15

	fmt.Println(longestSubarray(nums, k))
}

/*   Interview Tip
 This is a very important hashing pattern:
  - Prefix Sum + Hash Map
 Remember:
  - Required Prefix = Current Prefix Sum - K

And for the longest subarray:
  - Always store the first occurrence of each prefix sum.

Also note: if the array contains only positive numbers, a sliding-window approach can be used. But because this problem says integers and may contain negative numbers, Prefix Sum + Hash Map is the safe general solution.