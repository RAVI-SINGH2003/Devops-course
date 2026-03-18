package main

import "fmt"

// enums = type + const

// type OrderStatus int

// const (
// 	Received OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )


type OrderStatus string

const (
	Received  = "Received"
	Confirmed = "Confirmed"
	Prepared  = "Prepared"
	Delivered = "Delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status changed to:", status)
}

func main() {
	changeOrderStatus(Prepared)
}
