//Kth Missing Positive Number
/* Given a sorted array of unique positive integers arr, your task is to return the kᵗʰ missing positive number that is not present in arr.
   The array is guaranteed to be strictly increasing, and the missing numbers are those positive integers that do not appear in arr but would appear in a full sequence starting from 1.
*/

package main

import "fmt"

func findKthPositive(arr []int, k int) int {
	left, right := 0, len(arr)-1

	// Binary search for the position where missing count >= k
	for left <= right {
		mid := left + (right-left)/2
		missing := arr[mid] - (mid + 1)

		if missing < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// left + k gives the exact k-th missing positive number
	return left + k
}

func main() {
	arr := []int{2, 3, 4, 7, 11}
	k := 5

	result := findKthPositive(arr, k)
	fmt.Printf("The %d-th missing positive number is: %d\n", k, result) // Output: 9
}
