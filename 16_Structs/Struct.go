/* -> AA struct is a user-defined type that groups related data together,it like a blueprint for an object.
// in go place of classes use struct
//IMP:
 --Struct = collection of fields
 --Access with .
 --Methods add behavior
 --Pointer receivers modify original data
 --Struct tags help with JSON and APIs
*/

package main

import (
	"fmt"
	"time"
)

/*
// Order struct

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision

}

/*
func main() {
	myOrder := order{
		id:     "1",
		amount: 50.00,
		status: "received",
	}
	myOrder.createdAt = time.Now()
	// get any field
	fmt.Println(myOrder.status)

	fmt.Println("Order struct", myOrder)

}  */

/*
// we can use struct as a blueprint
// by using struct we can do object oriented programing in go.
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision

}

func main() {
	myOrder := order{
		id:     "1",
		amount: 50.00,
		status: "received",
	}
	myOrder.createdAt = time.Now()
	// get any field
	fmt.Println(myOrder.status)

	myOrder2 := order{
		id:        "2",
		amount:    150,
		status:    "delivered",
		createdAt: time.Now(),
	}

	myOrder.status = "paid"

	fmt.Println("Order struct", myOrder2)
	fmt.Println("Order struct", myOrder)

}  */

/*
// Receiver type  func (o order)
// If you modifing the value of struct  then use star * pointer (o *order) not needed for return the value.
func (o *order) changeStatus(status string) {
	o.status = status // struct automaticall do the work of derefrenceing.

}

// we can add as many method as wanted.needed
func (o order) getAmount() float32 { // here we are not modifiying ony getting the value so * not needed.
	return o.amount

}

func main() {
	// if you don't set any field, default value is zero value.
	//zero value ->  int - 0, sring "", bool - false
	myOrder := order{
		id:     "1",
		amount: 50.00,
		status: "received",
	}
	myOrder.changeStatus("confirmed")
	//fmt.Println(myOrder)
	fmt.Println(myOrder.getAmount())


}  */
/*
// constructer in go (their is no constructer in go we r using a hack by making funs together with new keyword. )

func newOrder(id string, amount float32, status string) *order {

	//iitial setup goes here...
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder //we do not do directly return struct, return pointer.

}
func (o *order) changeStatus(status string) {
	o.status = status // struct automaticall do the work of derefrenceing.

}

// we can add as many method as wanted.needed
func (o order) getAmount() float32 { // here we are not modifiying ony getting the value so * not needed.
	return o.amount

}

func main() {
	myOrder := newOrder("1", 30.50, "received")
	fmt.Println(myOrder.amount)

}  */

/* // if need to use structs once so we can user this inline struct.
func main() {

	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)

}  */

//Struct Embedding

type customer struct {
	name  string
	phone string
}

//composition

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer
}

func main() {
	newCustomer := customer{
		name:  "Don",
		phone: "72503547",
	}
	newOrder := order{
		id:       "1",
		amount:   30,
		status:   "received",
		customer: newCustomer,
	}
	newOrder.customer.name = "Devil" // for change the name of customer
	fmt.Println(newOrder)
}
