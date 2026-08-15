func subarraysWithXorK(nums []int, k int) int {
	n := len(nums)
	xr := 0
	mpp := make(map[int]int)
	mpp[0]++
	cnt := 0

	for i := 0; i < n; i++ {
		xr = xr ^ nums[i]
		x := xr ^ k
		cnt += mpp[x]
		mpp[xr]++
	}
	return cnt
}
