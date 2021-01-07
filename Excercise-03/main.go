package main

import "fmt"

func main() {
	colors := map[string]string{
		"Red":   "#ff000",
		"Green": "#4bf7",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for index, color := range c {
		fmt.Printf("Index %s and color %s \n", index, color)
	}
}
