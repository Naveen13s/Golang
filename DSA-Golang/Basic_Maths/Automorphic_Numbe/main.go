package main

import "fmt"

func isAutomorphic(n int) bool {
	square := n * n

	digits := 1
	temp := n

	for temp >= 10 {
		digits++
		temp /= 10
	}

	divisor := 1
	for i := 0; i < digits; i++ {
		divisor *= 10
	}

	return square%divisor == n
}

func main() {
	n := 25

	fmt.Println(isAutomorphic(n))
}
