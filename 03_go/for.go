package main

import "fmt"

func main() {

	for i := 0; i <= 3; i++ {

		if i == 2 {
			continue
		}

		fmt.Println(i)
	}

	for i := range 10 {
		fmt.Println(i)
	}
}
