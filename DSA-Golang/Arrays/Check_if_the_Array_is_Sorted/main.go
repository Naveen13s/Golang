package main

import "fmt"

func isSorted(arr []int, n int) bool {
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}

	return true
}

func main() {
	arr := []int{1, 2, 2, 4, 5}

	fmt.Println(isSorted(arr, len(arr)))
}
