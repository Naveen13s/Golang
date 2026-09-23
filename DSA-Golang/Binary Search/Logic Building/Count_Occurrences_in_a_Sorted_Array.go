//Count Occurrences in a Sorted Array
/*You are given a sorted array of integers arr and an integer target.
Your task is to determine how many times target appears in arr.
Return the count of occurrences of target in the array.
*/

package main

import "fmt"

func countOccurrences(arr []int, target int) int {
	first := findBound(arr, target, true)
	if first == -1 {
		return 0 // Target is not present in the array
	}
	last := findBound(arr, target, false)
	return last - first + 1
}

// Helper function to find either the first or last index of target
func findBound(arr []int, target int, isFirst bool) int {
	left, right := 0, len(arr)-1
	bound := -1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			bound = mid
			if isFirst {
				right = mid - 1 // Keep searching left for first index
			} else {
				left = mid + 1 // Keep searching right for last index
			}
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return bound
}

func main() {
	arr := []int{1, 2, 2, 2, 2, 3, 4, 7, 8, 8}

	fmt.Println(countOccurrences(arr, 2))
	fmt.Println(countOccurrences(arr, 8))
	fmt.Println(countOccurrences(arr, 5))
}
