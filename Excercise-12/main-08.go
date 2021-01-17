package main

import (
	"sync"
)

func fanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexStream := make(chan interface{})
	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexStream <- i:
			}
		}
	}
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}
	go func() {
		wg.Wait()
		close(multiplexStream)
	}()
	return multiplexStream
}

func main() {

}
