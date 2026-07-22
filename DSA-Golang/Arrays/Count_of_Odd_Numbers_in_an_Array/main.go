package main

import "fmt"

func countOdd(arr []int, n int) int {
	count := 0

	for i := 0; i < n; i++ {
		if arr[i]%2 != 0 {
			count++
		}
	}

	return count
}

func main() {
	arr := []int{4, 7, 9, 2, 6, 5}

	fmt.Println(countOdd(arr, len(arr)))
}
