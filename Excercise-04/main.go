package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func (eb englishBot) getGreeting() string {
	return "Hello!"
}

func (esb spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(eb bot) {
	fmt.Println(eb.getGreeting())
}

func main() {
	eb := englishBot{}
	printGreeting(eb)
	sb := spanishBot{}
	printGreeting(sb)

}
