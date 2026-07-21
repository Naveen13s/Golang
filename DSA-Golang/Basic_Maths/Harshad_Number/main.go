package main

import "fmt"

func isHarshad(n int) bool {
	if n <= 0 {
		return false
	}

	original := n
	sum := 0

	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return original%sum == 0
}

func main() {
	n := 18

	fmt.Println(isHarshad(n))
}
