// Valid Anagram
package main

import "fmt"

// IsAnagram returns true if t is an anagram of s.
// Supports all ASCII lowercase characters in O(n) time and O(1) space.
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var counts [26]int
	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram")) // Output: true
	fmt.Println(isAnagram("rat", "car"))         // Output: false
}
