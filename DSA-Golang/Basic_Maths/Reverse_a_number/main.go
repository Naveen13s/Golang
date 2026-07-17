//Reverse a number
//You are given an integer n. Return the integer formed by placing the digits of n in reverse order.

package main

import "fmt"

func reverseNumber(n int) int {
	reverse := 0

	for n > 0 {
		digit := n % 10
		reverse = reverse*10 + digit
		n /= 10
	}

	return reverse
}

func main() {
	n := 1234
	fmt.Println(reverseNumber(n))
}
