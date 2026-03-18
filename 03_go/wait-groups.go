package main

import (
	"fmt"
	"sync"
)

func printt(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(num)
}

func main() {
    var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go printt(i,&wg)
	}
	wg.Wait()

}
