package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter no of sides of the polygon")
	var sides int
	fmt.Scan(&sides)

	var output string

	switch sides {
	case 3:
		output = "triangle"
	case 4:
		output = "quadrilateral"
	case 5:
		output = "pentagon"
	default:
		output = "Invalid input"
	}
	fmt.Printf("Answer is %s\n", output)

	// multiple condition switch

	day := "Monday"

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's workday")
	}

	// type , interface {} is like any
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("It's an integer")
		case string:
			fmt.Println("It's a string")
		case bool:
			fmt.Println("it's a boolean")
		default:
			fmt.Println("other", t)
		}
	}
	whoAmI("golang")
}
