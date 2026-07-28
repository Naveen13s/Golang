package main

import "fmt"

func union(nums1, nums2 []int) []int {
	i, j := 0, 0
	result := []int{}

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			if len(result) == 0 || result[len(result)-1] != nums1[i] {
				result = append(result, nums1[i])
			}
			i++
		} else if nums1[i] > nums2[j] {
			if len(result) == 0 || result[len(result)-1] != nums2[j] {
				result = append(result, nums2[j])
			}
			j++
		} else {
			if len(result) == 0 || result[len(result)-1] != nums1[i] {
				result = append(result, nums1[i])
			}
			i++
			j++
		}
	}

	for i < len(nums1) {
		if len(result) == 0 || result[len(result)-1] != nums1[i] {
			result = append(result, nums1[i])
		}
		i++
	}

	for j < len(nums2) {
		if len(result) == 0 || result[len(result)-1] != nums2[j] {
			result = append(result, nums2[j])
		}
		j++
	}

	return result
}

func main() {
	nums1 := []int{1, 2, 2, 3, 5}
	nums2 := []int{2, 3, 4, 5, 6}

	fmt.Println(union(nums1, nums2))
}

/*This is a classic Two Pointer problem. The important observation is that both arrays are sorted, allowing me to solve it in O(N + M) time without using a hash set. This pattern is also used in problems like Merge Two Sorted Arrays, Intersection of Two Sorted Arrays, and Merge Intervals.*?