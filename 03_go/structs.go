package main

import (
	"fmt"
	"time"
)

// order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
}

func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func main() {
	//if you don't set value of any field then default value is used
	// int => 0 , float => 0 , string=>"", bool=> false
	myOrder := order{
		id:     "1",
		amount: 1000.12,
		status: "received",
	}
	myOrder.createdAt = time.Now()
	myOrder.changeStatus("confirmed")
	fmt.Println(myOrder, myOrder.amount)

	order2 := newOrder("2", 12.1, "sent")
	fmt.Println(order2)


	//inline structs
	student := struct {
		name  string
		marks float64
	}{"ravi singh", 124.1}

	fmt.Println(student)

}
