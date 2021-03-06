package main

import "fmt"

func main() {

	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 0; i < 10; i++ {
			intStream <- i
		}

	}()
	for integer := range intStream {
		fmt.Printf("%v\n", integer)
	}
}
