package main

import (
	"fmt"
	"time"
)

func print(num int) {
	fmt.Println(num)
}
func main() {
	for i := 0; i < 10; i++ {
		// go print(i)
		go func(i int){
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second * 2)
}
