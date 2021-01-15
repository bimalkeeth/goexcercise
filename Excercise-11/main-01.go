package main

import "fmt"

func main() {

	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello Channel"
		close(stringStream)
	}()

	salutation, ok := <-stringStream
	fmt.Printf("(%v):%v", ok, salutation)
}
