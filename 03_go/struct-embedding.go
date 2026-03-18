package main

import "fmt"

type CUSTOMER struct {
	name string
}

func (c *CUSTOMER) changeName(name string) {
	c.name = name
}

type ORDER struct {
	id       string
	amount   float64
	customer CUSTOMER
}

func (o *ORDER) changeAmount(amount float64) {
	o.amount = amount
}

func main() {
	newCustomer := CUSTOMER{
		name: "ravi singh",
	}

	newOrder := ORDER{
		id:       "1",
		amount:   12.1,
		customer: newCustomer,
	}

	newOrder.customer.changeName("hello")

	newOrder.changeAmount(10210.1)

	fmt.Println(newOrder)
}
