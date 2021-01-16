package main

import "fmt"

func Generator(done <-chan interface{}, integers ...int) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for _, i := range integers {
			select {
			case <-done:
				return
			case intStream <- i:
			}
		}
	}()
	return intStream
}
func multiply(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
	multipliedStream := make(chan int)
	go func() {
		defer close(multipliedStream)
		for i := range intStream {
			select {
			case <-done:
				return
			case multipliedStream <- i * multiplier:
			}
		}
	}()
	return multipliedStream

}
func adding(done <-chan interface{}, intstrem <-chan int, additive int) <-chan int {
	addStream := make(chan int)
	go func() {
		defer close(addStream)
		for i := range intstrem {
			select {
			case <-done:
				return
			case addStream <- i + additive:
			}
		}

	}()
	return addStream
}

func main() {

	done := make(chan interface{})
	defer close(done)

	intStream := Generator(done, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	pipeline := multiply(done, adding(done, multiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}
}
