package main

import (
	"fmt"
	"sync"
)

func main() {

	mypool := &sync.Pool{New: func() interface{} {

		fmt.Println("Creating new instance")
		return struct{}{}
	}}
	mypool.Get()
	instance := mypool.Get()
	mypool.Put(instance)
	mypool.Get()

}
