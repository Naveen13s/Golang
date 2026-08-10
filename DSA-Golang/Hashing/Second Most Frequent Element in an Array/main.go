package main

import (
	"fmt"
	"math"
)

func secondMostFrequent(nums []int) int {
	freq := make(map[int]int)

	// Count frequencies
	for _, num := range nums {
		freq[num]++
	}

	// Find maximum frequency
	maxFreq := 0
	for _, count := range freq {
		if count > maxFreq {
			maxFreq = count
		}
	}

	secondFreq := -1
	answer := math.MaxInt

	// Find second highest frequency
	for num, count := range freq {

		if count == maxFreq {
			continue
		}

		if count > secondFreq {
			secondFreq = count
			answer = num
		} else if count == secondFreq && num < answer {
			answer = num
		}
	}

	if secondFreq == -1 {
		return -1
	}

	return answer
}

func main() {
	nums := []int{1, 2, 2, 3, 3, 3}

	fmt.Println(secondMostFrequent(nums))
}

/*  **Interview Tip**

This problem is an extension of Most Frequent Element.

The key idea is to distinguish between:
 - Highest frequency
 - Second highest distinct frequency

Remember to handle these edge cases:
 - All elements have the same frequency → return -1.
 - Multiple elements share the second highest frequency → return the smallest element.
Using a Hash Map for frequency counting gives the optimal O(N) solution.  */
