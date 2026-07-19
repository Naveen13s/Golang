package main

import "fmt"

func gcd(n1, n2 int) int {
	for n2 != 0 {
		n1, n2 = n2, n1%n2
	}

	return n1
}

func main() {
	n1 := 48
	n2 := 18

	fmt.Println("GCD:", gcd(n1, n2))
}
