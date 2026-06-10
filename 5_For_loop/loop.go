package main

import "fmt"

// for -> only construct in go for looping

func main() {
	//while loop - in the no while keywords like others
	//i := 1
	//for i <= 3 {
	//fmt.Println(i)
	//i = i + 1

	//}

	// Infinite loop (ctrl + c for cancel)
	//for {
	//println("1")

	//Classic For loop

	//for i := 0; i <= 3; i++ {
	//break

	//if i == 2 {
	//continue
	//}
	//fmt.Println(i)

	//1.22 range in go
	for i := range 3 {
		fmt.Println(i)
	}
}
