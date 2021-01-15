package main

import (
	"fmt"
	"sync"
)

func main() {

	var numCalcsCreated int
	calcPool := &sync.Pool{New: func() interface{} {
		numCalcsCreated += 1
		mem := make([]byte, 1024)
		return &mem
	}}
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorker = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorker)

	for i := numWorker; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
		}()
	}
	wg.Wait()
	fmt.Printf("%d Calculators were created", numCalcsCreated)

}
