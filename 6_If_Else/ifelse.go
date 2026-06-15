package main

import "fmt"

func main() {

	//age := 10

	//if age >= 18 {
	//fmt.Println("Person is adult")
	//} else {
	//fmt.Println("Person is not a adult")

	//}

	//if age >= 18 {
	//	fmt.Println("Person is an adult")
	//} else if age >= 12 {
	//	fmt.Println("person is teemager")

	//} else {
	//	fmt.Println("person is a kid")
	//}

	//Logical Operator
	//var role = "admin"
	//var hasPermissions = false

	//if role == "admin" || hasPermissions {
	//if role == "admin" && hasPermissions {
	//fmt.Println("yes")
	//}

	// Also we can declare variable inside if construct

	if age := 20; age >= 18 {
		fmt.Println("Person is an adult")
	  } else if age >= 12 {
		fmt.Println("Person is a teenager")

	}

	//Go does not have ternary operator, we have to use normal if else
}
