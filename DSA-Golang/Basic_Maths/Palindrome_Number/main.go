//Palindrome Number
//You are given an integer n. You need to check whether the number is a palindrome number or not. Return true if it's a palindrome number, otherwise return false.
//A palindrome number is a number which reads the same both left to right and right to left.

package main

import "fmt"

func isPalindrome(n int) bool {
	original := n
	reverse := 0

	for n > 0 {
		digit := n % 10
		reverse = reverse*10 + digit
		n /= 10
	}

	return original == reverse
}

func main() {
	n := 121

	fmt.Println(isPalindrome(n))
}
