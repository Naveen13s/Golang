package main

import (
	"fmt"
	"math"
)

func mostFrequentElement(nums []int) int {
	frequency := make(map[int]int)

	for _, num := range nums {
		frequency[num]++
	}

	maxFreq := 0
	answer := math.MaxInt

	for num, freq := range frequency {
		if freq > maxFreq {
			maxFreq = freq
			answer = num
		} else if freq == maxFreq && num < answer {
			answer = num
		}
	}

	return answer
}

func main() {
	nums := []int{1, 2, 2, 3, 3}

	fmt.Println(mostFrequentElement(nums))
}

// Alternative (Without math.MaxInt)
func mostFrequentElement(nums []int) int {
	frequency := make(map[int]int)

	for _, num := range nums {
		frequency[num]++
	}

	answer := nums[0]
	maxFreq := 0

	for num, freq := range frequency {
		if freq > maxFreq || (freq == maxFreq && num < answer) {
			maxFreq = freq
			answer = num
		}
	}

	return answer
}

/*💡 Interview Tip
There are two common approaches:

Approach:- 	                                     Time	    Space
1. Brute Force (count each element separately)	     O(N²)	     O(1)
2. Hash Map (Frequency Counting)	                 O(N)	     O(N) ✅
 -> The Hash Map approach is the expected interview solution because it counts frequencies in a single traversal and easily handles the tie-breaking condition (choosing the smaller element when frequencies are equal).  */
