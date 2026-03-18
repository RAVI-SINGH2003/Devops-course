package main

import (
	"fmt"
)

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	//variadic function is that which can take n no of parameters.

	fmt.Println(1, 2, 3, 4, 5, 6, "hello")

	fmt.Println(sum(1, 2, 3, 4, 5))

	array := []int{10, 20, 30}

	fmt.Println(sum(array...))

}
