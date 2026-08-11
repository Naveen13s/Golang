package main

import (
	"fmt"
	"math"
)

func sumHighestAndLowestFrequency(arr []int) int {
	freq := make(map[int]int)

	// Count frequencies
	for _, num := range arr {
		freq[num]++
	}

	maxFreq := 0
	minFreq := math.MaxInt

	// Find highest and lowest frequency
	for _, count := range freq {
		if count > maxFreq {
			maxFreq = count
		}

		if count < minFreq {
			minFreq = count
		}
	}

	return maxFreq + minFreq
}

func main() {
	arr := []int{1, 2, 2, 3, 3, 3}

	fmt.Println(sumHighestAndLowestFrequency(arr))
}
