package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	set := make(map[int]bool)

	for _, num := range nums {
		set[num] = true
	}

	longest := 0

	for _, num := range nums {

		// Start only if previous number doesn't exist
		if !set[num-1] {

			current := num
			length := 1

			for set[current+1] {
				current++
				length++
			}

			if length > longest {
				longest = length
			}
		}
	}

	return longest
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}

	fmt.Println(longestConsecutive(nums))
}

/* Interview Tip:

There are three common approaches:

Approach	    Time	     Space
Brute Force   	O(N²)	      O(1)
Sorting	      O(N log N)	O(1) (excluding sorting)
Hash Set	    O(N)	      O(N) ✅
-> The Hash Set approach is the expected interview solution because it avoids sorting and processes each sequence only once by starting from numbers that have no predecessor.
-> Hashing Pattern Learned: This problem introduces an important hashing technique—using a Hash Set for fast lookups (O(1)) rather than a frequency map. The key idea is to identify the start of a sequence (num - 1 doesn't exist) and then expand it efficiently. This pattern appears in many interview problems involving consecutive elements or membership checks.
*/
