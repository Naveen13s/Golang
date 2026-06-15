//closure is a function that remembers and can access variables from outside its own scope, even after the outer function has finished executing.
//Closures allow functions to carry their own data/state without using global variables

package main

import "fmt"

func counter() func() int {

	var count int = 0
//captures variables from outer scope
	return func() int {
		count += 1
		return count
	}

}
func main() {
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())

}
