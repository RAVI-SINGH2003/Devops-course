package main

import "fmt"

func print1[T comparable, V comparable](array []T, val V) {
	for _, e := range array {
		fmt.Println(e)
	}
	fmt.Println(val)
}

// func print2[T int | string | bool, V comparable](array []T, val V) {
// 	for _, e := range array {
// 		fmt.Println(e)
// 	}
// 	fmt.Println(val)
// }

type Student[T any] struct {
	marks []T
}

func main() {
	print1([]int{1, 2, 3}, "hello")
	student1 := Student[int]{
		marks: []int{1, 2, 3},
	}
	fmt.Println(student1)
}
