package main

import "fmt"

func main() {
	chanOwner := func() <-chan int {
		result := make(chan int, 5)
		go func() {
			defer close(result)
			for i := 0; i <= 5; i++ {
				result <- i
			}
		}()
		return result
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received : %d\n", result)
		}
		fmt.Println("All received!")
	}
	result := chanOwner()
	consumer(result)

}
