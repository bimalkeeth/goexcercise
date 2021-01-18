package main

import (
	"fmt"
	"sync"
	"time"
)

type job struct {
	i    int
	max  int
	text string
}

func outputText(j *job, wg *sync.WaitGroup) {
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(j.text)
		j.i++
	}
	wg.Done()
}

func main() {
	wg := new(sync.WaitGroup)
	fmt.Println("Starting")

	hello := new(job)
	hello.text = "Hello"
	hello.i = 0
	hello.max = 2

	world := new(job)
	world.text = "world"
	world.i = 0
	world.max = 2

	go outputText(hello, wg)
	go outputText(world, wg)

	wg.Add(2)
	wg.Wait()
}
