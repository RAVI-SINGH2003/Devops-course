package main

import (
	"fmt"
	"slices"
)

func main() {
	//slices - dynamic arrays

	// uninitialized slice is nil
	var nums []int
	fmt.Println(nums, nums == nil, nil, len(nums))

	// make(type,size,capacity)
	// initially capacity = size
	// both capacity and size keeps on changing
	// capacity becomes double as needed
	var nums1 = make([]int, 2, 2)
	// var nums1 = make([]int, 0, 2)

	fmt.Println(nums1, len(nums1), cap(nums1))
	nums1 = append(nums1, 2)
	nums1 = append(nums1, 3)
	nums1 = append(nums1, 4)
	nums1 = append(nums1, 5)
	nums1 = append(nums1, 6)
	nums1[0] = 1
	fmt.Println("After appending", nums1, len(nums1), cap(nums1))

	// shorthand
	nums2 := []int{1, 2}
	nums2 = append(nums2, 3)
	fmt.Println(nums2)

	//copy
	var nums3 = make([]int, len(nums2))
	copy(nums3, nums2)
	fmt.Println(nums3)

	// slice operator [startIndex: endIndex], endIndex is not included
	var nums5 = []int{1, 2, 3}
	fmt.Println(nums5[0:3], nums5[:3], nums5[1:], nums5[:])

	//slice pkg
	var nums6 = []int{1, 2}
	var nums7 = []int{1, 2}
	var nums8 = make([]int, 0)
	nums8 = append(nums8, 1)
	nums8 = append(nums8, 2)
	fmt.Println(slices.Equal(nums6, nums7), slices.Equal(nums6, nums8))

	//2D slices

	var nums9 = [][]int{{1, 2, 3}, {4, 5, 6}}
	var nums10 [][]int
	tempp := []int{1, 2, 3}
	nums10 = append(nums10, tempp)
	fmt.Println(nums9, nums10)

}
