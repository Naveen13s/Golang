package main

import (
	"fmt"
	"maps"
)

//maps- Associative DS -> hash , object, dict as python

func main() {
	//creating map
	/*
		m := make(map[string]string)

		//setting an element

		m["name"] = "golang"
		m["area"] = "backend"

		//get an element
		fmt.Println(m["name"], m["area"])

		// IMP: if key does not exit in the map then it returns zero
		//fmt.Println(m["naveen"])  */
	/* - > if value int it will return 0

	m := make(map[string]int)
	m["age"] = 30
	fmt.Println(m["phone"]) */

	/*   Length of Map
	m := make(map[string]int)
	m["age"] = 30
	m["price"] = 300
	fmt.Println(len(m))   */

	/*  Delete Fun
	m := make(map[string]int)
	m["age"] = 30
	m["price"] = 300

	delete(m, "price")
	fmt.Println(m)    */

	/*  Clear fun
	m := make(map[string]int)
	m["age"] = 30
	m["price"] = 300
	clear(m)
	fmt.Println(m)   */

	/*  Map withou make fun (If you know the element from starting use this otherwise use make fun)

	m := map[string]int{"price": 20, "phone": 2}
	fmt.Println(m)    */

	// To check weather our elemnet in under map or not

	m := map[string]int{"price": 20, "phone": 2}

	v, ok := m["price"]
	fmt.Println(v)

	if ok {
		fmt.Println("All ok")

	} else {
		fmt.Println("not Ok")
	}

	// to compare a map use maps package
	m1 := map[string]int{"price": 20, "phone": 2}
	m2 := map[string]int{"price": 20, "phone": 2}
	fmt.Println(maps.Equal(m1, m2))

}
