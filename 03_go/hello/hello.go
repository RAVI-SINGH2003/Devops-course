package main

import (
	"fmt"
	"go/greetings"
)

func main() {
	fmt.Println("Enter your name:")
	var name string
	fmt.Scan(&name)
	message := greetings.Hello(name)
	fmt.Println(message)
}
