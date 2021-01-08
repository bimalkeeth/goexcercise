package main

import "fmt"

func main() {

	writeSomething()
	fmt.Println("shutting down")
}

func writeSomething() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error occurred")
		}
	}()

	panic("write operation error")
}
