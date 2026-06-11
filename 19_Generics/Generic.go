package main

import "fmt"

/*
// approx same function is duplication so to solve this problems we will use Generics.
func printSlice(items []int) {

	for _, item := range items {
		fmt.Println(item)

	}
}

func printStringSlice(items []string) {

	for _, item := range items {
		fmt.Println(item)

	}
}

func main() {
	//nums := ([]int{1, 2, 3})
	names := []string{"golang", "typesscript"}
	printStringSlice(names)
	//printSlice()

}  */

/*
//Generics on the function
//func printSlice[T any](items []T) { // we can also use [T interface{}]
//func printSlice[T int | string | bool](items []T) {
//func printSlice[T comparable , V string](items []T, names v) { // we can also use multiple generics types
func printSlice[T comparable](items []T) { // comparable used for all types.
	for _, item := range items {
		fmt.Println(item)

	}
}

func main() {
	//nums := ([]int{1, 2, 3})
	//names := []string{"golang", "typesscript"}
	vals := []bool{true, false, true}
	//printStringSlice(names)
	//printSlice(nums)
	//printSlice(names)
	printSlice(vals)

}  */

// Stucts
// Stack - LIFO
/*
type stack struct { // for integer on;ly
	elements []int
}

func main() {

	myStack := stack{

		elements: []int{1, 2, 3},
	}
	fmt.Println(myStack)

} */

//string

type stack[T any] struct {
	elements []T
}

func main() {

	myStack := stack[string]{

		elements: []string{"golang"},
	}
	fmt.Println(myStack)

}
