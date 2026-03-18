package main

import "fmt"

func main() {

	// printing slices/arrays
	nums := []int{1, 2, 3}

	nums1 := [2]int{100, 200}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	for i := range len(nums) {
		fmt.Println(nums[i])
	}

	for i, num := range nums {
		fmt.Println(num, i)
	}
	for _, num := range nums1 {
		fmt.Println(num)
	}

	//printing maps
	m := map[string]string{"name": "ravi singh", "address": "New Delhi"}

	for k, v := range m {
		fmt.Println(k, v)
	}

	//c is actually ascii/unicode value
	//unicode code point rune
	// i is starting byte of rune
	//255 -> 1byte, 2 byte
	for i, c := range "golang" {
		fmt.Println(i, c, string(c))
	}

}
