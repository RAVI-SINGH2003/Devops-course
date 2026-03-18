package main

import (
	"fmt"
	"maps"
)

//maps: = hashtables, object , dict
// if key does not exist in map default values of datatypes are returned

func main() {
	mpp := make(map[string]string)
	mpp["name"] = "Ravi Singh"
	mpp["age"] = "twenty two"
	fmt.Println(mpp, mpp["name"], mpp["phone"])

	mpp1 := make(map[string]int)
	mpp1["age"] = 22
	mpp1["tax"] = 100
	fmt.Println(mpp1["age"], mpp1["aed"])
	fmt.Print(len(mpp1))
	delete(mpp1, "age")
	// clear(mpp1)
	fmt.Println(mpp1)

	// making map without make
	mpp3 := map[string]int{"price": 1, "marks": 90}
	mpp3["age"] = 80
	fmt.Println(mpp3)

	value, ok := mpp3["e"]

	if ok {
		fmt.Println("all ok", value)
	} else {
		fmt.Println("element does not exist")
	}

	mpp4 := map[string]int{"name": 1}
	mpp5 := map[string]int{"name": 5}
	fmt.Println(maps.Equal(mpp4, mpp5))
}
