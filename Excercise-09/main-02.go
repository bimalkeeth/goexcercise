package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	fmt.Println("Initializing go application")
}
func main() {

	go func() {
		for range time.Tick(time.Second * 2) {
			fmt.Println("Engine #2 is working ")
		}
	}()

	for range time.Tick(time.Second * 2) {
		fmt.Println("Engine #1 is working ", runtime.NumGoroutine(), " tasks(Go Routines) running")
	}

}
