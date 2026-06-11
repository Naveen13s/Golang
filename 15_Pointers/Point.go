// A pointer stores the memory address of another variable
/*Pointers are useful when you want to:
->modify a variable inside a function
->avoid copying large data structures
->share data efficiently
//Symbols:
 ->  &x        -     address of x
 ->  *p        -     value stored at pointer p
 ->  var p *int  -   pointer to int

*/

package main

import "fmt"

//by value
//func changeNum(num int) {
//num = 5
//fmt.Println("In changeNum", num)
//}

//by refrence
func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum", *num)
}

func main() {
	num := 1
	//changeNum(num)
	changeNum(&num)

	// for memeory address
	//fmt.Println("Memory address", &num)
	fmt.Println("After changeNum is main", num)

}
