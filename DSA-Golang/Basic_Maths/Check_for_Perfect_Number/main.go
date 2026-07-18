package main

import "fmt"

func isPerfect(n int) bool {
	if n <= 1 {
		return false
	}

	sum := 1

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i

			if i != n/i {
				sum += n / i
			}
		}
	}

	return sum == n
}

func main() {
	n := 28

	fmt.Println(isPerfect(n))
}
