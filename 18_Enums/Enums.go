package main

import "fmt"

//enumerated types

type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

/*
type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
) */

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)

}

/*
func changeOrderStatus(status string) {
	fmt.Println("changing order status to", status)

} */
func main() {
	//changeOrderStatus("received")
	changeOrderStatus(Received)
	//changeOrderStatus(Prepared)

}
