package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "Hello from channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- "Hello from channel 2"
	}()

	for {
		select {
		case chan1 := <-channel1:
			fmt.Println(chan1)
		case chan2 := <-channel2:
			fmt.Println(chan2)
		case <-time.After(time.Second * 1):
			return
		}
	}

}
