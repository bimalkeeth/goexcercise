package main

import (
	"fmt"
	"sync"
)

var doOnce sync.Once

func runMe() {
	doOnce.Do(func() {
		fmt.Println("I have been run once")
	})
	fmt.Println("I have been run")
}

func main() {
	runMe()
	runMe()
}
