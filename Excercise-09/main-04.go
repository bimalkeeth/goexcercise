package main

import (
	"fmt"
	"sync"
	"time"
)

type data struct {
	lock  *sync.RWMutex
	round map[string]int
}

func newData() *data {
	return &data{round: make(map[string]int), lock: &sync.RWMutex{}}
}

func (d *data) update(wid string) {
	d.lock.Lock()
	defer d.lock.Unlock()
	count, ok := d.round[wid]
	if !ok {
		fmt.Println("some error occurred")
		return
	}
	d.round[wid] = count + 1
}
func doWork(wid string, d *data, wg *sync.WaitGroup) {
	for range time.Tick(time.Second * 2) {
		d.update(wid)
	}
	wg.Done()
}
func getData(d *data) {
	for range time.Tick(time.Second * 2) {
		d.lock.RLock()
		fmt.Println(d)
		d.lock.RUnlock()
	}
}

func main() {
	var wg sync.WaitGroup
	d := newData()
	d.round["one"] = 0
	d.round["two"] = 0
	d.round["three"] = 0
	go doWork("one", d, &wg)
	wg.Add(1)
	go doWork("two", d, &wg)
	wg.Add(1)
	go doWork("three", d, &wg)
	wg.Add(1)

	go getData(d)

	wg.Wait()
	fmt.Println("we will not get here")

}
