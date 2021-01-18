package main

import (
	"fmt"
	"sync"
	"time"
)

func printGreeting2(done <-chan interface{}) error {
	greeting, err := genGreeting2(done)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}
func printFarewell2(done <-chan interface{}) error {
	farewell, err := genFarewell2(done)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}
func genGreeting2(done <-chan interface{}) (string, error) {
	switch locale, err := locale2(done); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")

}

func genFarewell2(done <-chan interface{}) (string, error) {
	switch locale, err := locale2(done); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}
func locale2(done <-chan interface{}) (string, error) {
	select {
	case <-done:
		return "", fmt.Errorf("canceled")
	case <-time.After(1 * time.Minute):
	}
	return "EN/US", nil
}

func main() {
	var wg sync.WaitGroup
	done := make(chan interface{})
	defer close(done)
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreeting2(done); err != nil {
			fmt.Printf("%v", err)
			return
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell2(done); err != nil {
			fmt.Printf("%v", err)
			return
		}
	}()
	wg.Wait()
}
