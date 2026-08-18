// Largest Odd Number in a String
package main

import "fmt"

func largestOddNumber(s string) string {
	// Find the rightmost odd digit
	end := -1

	for i := len(s) - 1; i >= 0; i-- {
		digit := s[i] - '0'

		if digit%2 == 1 {
			end = i
			break
		}
	}

	// No odd digit found
	if end == -1 {
		return ""
	}

	// Remove leading zeros
	start := 0
	for start <= end && s[start] == '0' {
		start++
	}

	// If the substring contains only zeros
	if start > end {
		return ""
	}

	return s[start : end+1]
}

func main() {
	s := "000123450"

	result := largestOddNumber(s)

	fmt.Println(result)
}
