// Sort Characters by Frequency
package main

import (
	"fmt"
	"sort"
)

func frequencySort(s string) []string {
	// Step 1: Count frequencies of each character
	freqMap := make(map[rune]int)
	for _, char := range s {
		freqMap[char]++
	}

	// Step 2: Extract unique characters
	uniqueChars := make([]rune, 0, len(freqMap))
	for char := range freqMap {
		uniqueChars = append(uniqueChars, char)
	}

	// Step 3: Sort using custom comparator
	sort.Slice(uniqueChars, func(i, j int) bool {
		charA, charB := uniqueChars[i], uniqueChars[j]

		// Priority 1: Higher frequency first
		if freqMap[charA] != freqMap[charB] {
			return freqMap[charA] > freqMap[charB]
		}

		// Priority 2: Alphabetical order (ascending) for ties
		return charA < charB
	})

	// Step 4: Convert runes to string slice output
	result := make([]string, len(uniqueChars))
	for i, char := range uniqueChars {
		result[i] = string(char)
	}

	return result
}

func main() {
	fmt.Println(frequencySort("tree"))   // Output: ["e", "r", "t"]
	fmt.Println(frequencySort("cccaaa")) // Output: ["a", "c"]
	fmt.Println(frequencySort("Aabb"))   // Output: ["b", "A", "a"]
}
