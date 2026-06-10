package main

import "fmt"

// Iterating over data structure

func main() {
	//nums := []int{5, 6, 7}
	/* using for loop
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}   */

	/* sum

	sum := 0
	for _, num := range nums {
		sum = sum + num
	}
	fmt.Println(sum)  */

	/* This for slice if u have range so by using map we can Iterate

	for i, num := range nums {
		fmt.Println(num, i)
	}   */

	m := map[string]string{"Fname": "Naveen", "lname": "Kumar"}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// if only key is needed
	for k := range m {
		fmt.Println(k)
	}

	// Unicode code  point rune  (Range on string)
	//Starting byte of rune
	//255 -> 1 byte

	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}

}
