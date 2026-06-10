package main

import "fmt"

func main() {

	/* Simple switch

	i := 3

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")

	case 3:
		fmt.Println("three")

	default:
		fmt.Println("others")

	} */

	/*Multiple Condition Switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")

	default:
		fmt.Println("It's workday")
	}  */

	// Type Switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("It's an integer")

		case string:
			fmt.Println("It's string")

		case bool:
			fmt.Println("It's a boolean")

		default:
			fmt.Println("Other", t)

		}
	}

	whoAmI("Sun")
}
