package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"http://google.com", "http://facebook.com", "http://stackoverflow.com", "http://golang.org", "http://amazon.com"}

	c := make(chan string)

	for _, val := range links {
		go checkLink(val, c)
	}
	for l := range c {
		go func(x string) {
			time.Sleep(10 * time.Millisecond)
			checkLink(x, c)
		}(l)
	}

}
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
