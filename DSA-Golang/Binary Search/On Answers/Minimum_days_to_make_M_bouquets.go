//Minimum days to make M bouquets

func roseGarden(n int, nums []int, k int, m int) int {
	// If we don't have enough total roses, it's impossible
	if n < m*k {
		return -1
	}

	left, right := nums[0], nums[0]
	for _, day := range nums {
		if day < left {
			left = day
		}
		if day > right {
			right = day
		}
	}
	ans := -1

	for left <= right {
		mid := left + (right-left)/2

		// Count how many bouquets can be made by day 'mid'
		bouquets := 0
		consecutive := 0
		for _, day := range nums {
			if day <= mid {
				consecutive++
				if consecutive == k {
					bouquets++
					consecutive = 0
				}
			} else {
				consecutive = 0
			}
		}

		// If we can make enough bouquets, try finding an earlier day
		if bouquets >= m {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}