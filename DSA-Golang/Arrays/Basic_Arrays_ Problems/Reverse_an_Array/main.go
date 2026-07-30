package main

import "fmt"

func reverseArray(arr []int, n int) []int {
	left := 0
	right := n - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}

	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(reverseArray(arr, len(arr)))
}
