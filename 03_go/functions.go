package main

import "fmt"

func addNumbers(a int, b int) int {
	return a + b
}

func subNumbers(a, b int) int {
	return a - b
}

func addSliceValues(slice []int) int {
	sum := 0
	for _, num := range slice {
		sum += num
	}
	return sum
}

func getLanguages() (string, string, string) {
	return "golang", "javascript", "c++"
}

func process(fn func(a int) int) {
	fmt.Println(fn(2))
}

func returnFunc() func(a int) int {
	newFun := func(a int) int {
		return a * 2
	}
	return newFun
}

// functions are first class citizens means , they can be 
// passed ,returned to a function
func main() {
	fmt.Println(addNumbers(1, 2))
	fmt.Println(subNumbers(1, 2))
	fmt.Println(addSliceValues([]int{1, 2, 3}))

	lang1, lang2, _ := getLanguages()
	fmt.Println(lang1, lang2)

	// passing function to a function
	fn := func(a int) int {
		return a * a
	}
	process(fn)

	// returning a function to a function
	value := returnFunc()(2)
	fmt.Println("returned function value", value)

}
