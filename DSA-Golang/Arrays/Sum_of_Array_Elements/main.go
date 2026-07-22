package main

import "fmt"

func arraySum(arr []int) int {
	sum := 0

	for _, num := range arr {
		sum += num
	}

	return sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(arraySum(arr))
}
