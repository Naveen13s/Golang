//Search in rotated sorted array-II

func searchInARotatedSortedArrayII(nums []int, k int) bool {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == k {
			return true
		}
		// Handle duplicate edge case: shrink bounds when values are ambiguous
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
			continue
		}
		// Left side is sorted
		if nums[left] <= nums[mid] {
			if nums[left] <= k && k < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // Right side is sorted
			if nums[mid] < k && k <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}