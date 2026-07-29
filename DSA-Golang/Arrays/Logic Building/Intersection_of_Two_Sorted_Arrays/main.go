package main

import "fmt"

func intersection(nums1, nums2 []int) []int {
	i, j := 0, 0
	result := []int{}

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return result
}

func main() {
	nums1 := []int{1, 2, 2, 3, 4}
	nums2 := []int{2, 2, 4, 6}

	fmt.Println(intersection(nums1, nums2))
}
