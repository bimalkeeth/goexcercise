package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {

		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return

				}
			}
		}()
		return completed
	}
	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Cancelling doWork routine...")
		close(done)

	}()
	<-terminated
	fmt.Println("Done")
	// Perhaps more work is done here fmt.Println("Done.")
}
