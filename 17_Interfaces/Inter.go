//An interface defines a set of method signatures. Any type that implements those methods automatically satisfies the interface.
// //In Go, interfaces are implemented implicitly—you don't need an implements keyword.
/* IMP:
-> Interface defines behavior
-> Any type implementing methods satisfies it automatically
-> Enables polymorphism
-> interface{} can hold any value
-> Use type assertions and type switches to inspect underlying values.
-> Make testing easier (mocking)
-> Decouple implementation from usage
*/

package main

import "fmt"

//Interface have naming convenging

type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	gateway paymenter
	//gateway stripe
	//gateway razorpay  //concrit implement type should  be not used we need to provide interface.
}

// Open close principle
func (p payment) makePayment(amount float32) {
	//razorpayPaymnetGw := razorpay{}
	//razorpayPaymnetGw.pay(amount)
	//stripePaymentGW := stripe{}
	p.gateway.pay(amount)

}

type razorpay struct{}

//method signature should be same
func (r razorpay) pay(amount float32) {
	//logic to make payment
	fmt.Println("Making payment using razorpay", amount)

}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Making payment using stripe", amount)

}

// unit testing
type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("making payment using fake gateway for testig purpose")
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("Making payment using paypal", amount)
}

func (p paypal) refund(amount float32, account string) {

}

func main() {
	//stripePaymentGW := stripe{}
	//razorpayPaymentGW := razorpay{}
	//fakeGW := fakepayment{}
	paypalGW := paypal{}
	newPayment := payment{
		gateway: paypalGW,
		//gateway: fakeGW,
		//gateway: razorpayPaymentGW,
	}
	newPayment.makePayment(100)

}
