package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

type razorpay struct{}
func (r razorpay) pay(amount float32) {
	fmt.Println("Razorpay payment of: ", amount)
}

type stripe struct{}
func (r stripe) pay(amount float32) {
	fmt.Println("Stripe payment of: ", amount)
}

func main() {
	razorpayGw := razorpay{}
	stripeGw := stripe{}

	payment1 := payment{
		gateway: stripeGw,
	}
	payment2 := payment{
		gateway: razorpayGw,
	}
	payment1.gateway.pay(200)
	payment2.gateway.pay(300)
}
