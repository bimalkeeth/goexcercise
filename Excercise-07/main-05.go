package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("Count of iteration ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()

}
