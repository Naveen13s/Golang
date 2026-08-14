package main

import "fmt"

func subarraySum(nums []int, k int) int {
	prefixSum := 0
	count := 0

	// prefixSum -> frequency
	prefixMap := make(map[int]int)

	// Handles subarrays starting from index 0
	prefixMap[0] = 1

	for _, num := range nums {
		prefixSum += num

		// Check how many previous prefix sums
		// can form sum k
		count += prefixMap[prefixSum-k]

		// Store current prefix sum
		prefixMap[prefixSum]++
	}

	return count
}

func main() {
	nums := []int{1, 2, 3}
	k := 3

	fmt.Println(subarraySum(nums, k))
}

/* Interview Tip

Remember the difference between these two hashing problems:

| Problem                        | What to store in Hash Map     |
| ------------------------------ | ----------------------------- |
| Longest Subarray With Sum K    | `prefixSum → first index`     |
| **Count Subarrays With Sum K** | **`prefixSum → frequency`** ✅ |

The core formula remains the same:
 - previousPrefixSum = currentPrefixSum - K
But the thing stored in the map changes according to the question. */
