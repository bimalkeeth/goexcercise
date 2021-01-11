package main

//Dead lock generation

import (
	"fmt"
	"sync"
	"time"
)

type Value struct {
	mu    sync.Mutex
	Value int
}

var wg sync.WaitGroup

func PrintSum(v1, v2 *Value) {
	defer wg.Done()
	v1.mu.Lock()
	defer v1.mu.Unlock()

	time.Sleep(2 * time.Second)

	v2.mu.Lock()
	defer v2.mu.Unlock()

	fmt.Printf("sum %v\n", v1.Value+v2.Value)
}

func main() {
	var a, b Value

	a = Value{Value: 5}
	b = Value{Value: 6}
	wg.Add(2)
	go PrintSum(&a, &b)
	go PrintSum(&b, &a)
	wg.Wait()
}
