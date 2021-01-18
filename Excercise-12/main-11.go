package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan interface{}, pulseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {
	heartbeat := make(chan interface{})
	result := make(chan time.Time)

	go func() {
		defer close(heartbeat)
		defer close(result)

		pulse := time.Tick(pulseInterval)
		workGen := time.Tick(2 * pulseInterval)

		sendPulse := func() {
			select {
			case heartbeat <- struct{}{}:
			default:
			}
		}
		sendResult := func(r time.Time) {
			for {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case result <- r:
					return
				}
			}
		}

		for {
			select {
			case <-done:
				return
			case <-pulse:
				sendPulse()
			case r := <-workGen:
				sendResult(r)
			}
		}

	}()
	return heartbeat, result
}

func main() {

	done := make(chan interface{})
	time.AfterFunc(1*time.Minute, func() { close(done) })
	const timeout = 2 * time.Second
	heartbeat, result := doWork(done, timeout/2)
	for {
		select {
		case _, ok := <-heartbeat:
			if ok == false {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-result:
			if ok == false {
				return
			}
			fmt.Printf("Result %v\n", r.Second())
		case <-time.After(timeout):
			return
		}
	}

}
