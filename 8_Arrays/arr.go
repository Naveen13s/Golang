package main

import "fmt"

// Number sequence of specific language
func main() {
	/*zero values in int case
	  var nums [4]int

	  nums[0] = 1
	  fmt.Println(nums[0])
	   to print whole Arrays
	   fmt.Println(nums) */

	/* to get to know array length
	fmt.Println(len(nums))  */

	/*false in bool
	var vals [4]bool
	vals[2] = true
	fmt.Println(vals)  */

	/* string
	var name [3]string
	name[0] = "golang"
	fmt.Println(name) */

	/* To declare in single line
	nums := [3]int{1, 2, 3}

	fmt.Println(nums)  */

	//2d Arrays
	nums := [2][2]int{{2, 5}, {6, 8}}
	fmt.Println(nums)
	//key Point:
	// fixed size. that is predictable
	// Memory optimization
	// contant time access

}
