//Floor and Ceil in Sorted Array
/* Given a sorted array nums and an integer x. Find the floor and ceil of x in nums. The floor of x is the largest element in the array which is smaller than or equal to x. The ceiling of x is the smallest element in the array greater than or equal to x. If no floor or ceil exists, output -1.*/

func getFloorAndCeil(nums []int, x int) []int {
	floor := -1
	ceil := -1

	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == x {
			// Exact match is both the floor and the ceil
			return []int{nums[mid], nums[mid]}
		} else if nums[mid] < x {
			// Candidate for floor, look right for a larger possible floor
			floor = nums[mid]
			low = mid + 1
		} else {
			// Candidate for ceil, look left for a smaller possible ceil
			ceil = nums[mid]
			high = mid - 1
		}
	}

	return []int{floor, ceil}
}
