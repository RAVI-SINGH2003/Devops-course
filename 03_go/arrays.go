package main

import "fmt"

func main() {
	//arrays
	var nums [20]string
	fmt.Println(nums,len(nums))
	nums[2] = "hello"
	fmt.Println(nums)

	// declaration and assignment
	nums1 := [3]int{1, 2, 3}
	fmt.Println(nums1)

	//2d arrays
	nums2 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(nums2)
	
}
