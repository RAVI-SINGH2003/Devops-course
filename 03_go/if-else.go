package main

import "fmt"

func main() {
	if age := 6; age >= 11 {
		fmt.Println(age)
	} else if age < 5 {
		fmt.Println(age)
	} else {
		fmt.Println("s", age)
	}
}
