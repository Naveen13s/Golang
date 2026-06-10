package main

import "fmt"

/*
func add(a int, b int) int {

	return a + b

}
	func main() {
	result := add(3, 5)
	fmt.Print(result)

}    */

func getLanguages() (string, string, string) { // we print multiple value but need give type here and type shoud be in same order
	return "golang", "javascript", "c"
}

func main() {
	a1, a2, a3 := getLanguages()
	fmt.Println(a1, a2, a3)
}

/*  func main() {  //  this also use
	fmt.Println(getLanguages())

} */

func processIt() func(a int) int {
	//fn(1)
	return func(a int) int {
		return 4
	}
}

func main() {
	fn := processIt()
	fn(6)
	//fn := func (a int)int{
	//return 2
}
