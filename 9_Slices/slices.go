package main

import "fmt"

//slice -> Dynamic
//Most used construct in go
// + useful method

func main() {

	/*uninitalized slice is nil
	var nums []int
	//fmt.Println(nums == nil)
	fmt.Println(len(nums))  */

	/* var nums = make([]int, 2,)
	//2. Capacity -> maximum number of elements can fit
	fmt.Println(cap(nums))
	//fmt.Println(nums == nil) */
	// now capicity is 5
	/* var nums = make([]int, 0, 5)
	//fmt.Println(cap(nums))
	//now capacity is 20 because slice is dynamic.
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)

	fmt.Println(nums)
	fmt.Println(cap(nums))  */

	/*ohter
	nums := []int{} //< capicity is zero> before append
	nums = append(nums, 1)
	nums = append(nums, 3)
	fmt.Println(nums)
	fmt.Println(cap(nums))
	fmt.Println(len(nums))   */

	/* We can use index if we want to add element with the help of iindex.
	if we have to add element on zero and 1st.
	var nums = make([]int, 2, 5)
	nums[0] = 4
	nums[1] = 6

	fmt.Println(nums)
	fmt.Println(cap(nums))
	fmt.Println(len(nums))   */

	/* copy Function

	var nums = make([]int, 0, 5)
	nums = append(nums, 2)
	var nums2 = make([]int, len(nums))

	copy(nums2, nums)
	fmt.Println(nums, nums2)  */

	/*  // Slice Operator

	var nums = []int{1, 2, 3, 4, 5, 6}

	//fmt.Println(nums[0:2])
	//fmt.Println(nums[:1])
	fmt.Println(nums[1:])   */

	/*  Slices Package - comparing element in increasing fashion -> so same it will return true otherwise false

	//var nums1 = []int{1, 2}
	//var nums2 = []int{1, 2}

	var nums1 = []int{1, 2, 4, 5}
	var nums2 = []int{1, 2, 4, 6}
	fmt.Print(slices.Equal(nums1, nums2))   */

	// 2-D Slices
	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums)

}
