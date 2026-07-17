package main

import "fmt"

func largestDigit(n int) int {
	if n == 0 {
		return 0
	}

	maxDigit := 0

	for n > 0 {
		digit := n % 10

		if digit > maxDigit {
			maxDigit = digit
		}

		n /= 10
	}

	return maxDigit
}

func main() {
	n := 5832

	fmt.Println("Largest Digit:", largestDigit(n))
}
