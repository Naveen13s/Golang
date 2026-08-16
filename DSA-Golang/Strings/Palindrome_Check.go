//Palindrome Check
/* You are given a string s. Return true if the string is palindrome, otherwise false.
A string is called palindrome if it reads the same forward and backward.
*/

package main

import (
	"fmt"
)

// Function to check if a given string is a palindrome
func palindromeCheck(s string) bool {
	left := 0
	right := len(s) - 1

	// Iterate while start pointer is less than end pointer
	for left < right {
		// If characters don't match, it's not a palindrome
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	str := "racecar"

	if palindromeCheck(str) {
		fmt.Printf("%s is a palindrome.\n", str)
	} else {
		fmt.Printf("%s is not a palindrome.\n", str)
	}
}
