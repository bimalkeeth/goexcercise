package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(ReadLine(5))
}

func ReadLine(lineNumber int) string {
	file, _ := os.Open("Excercise-07/ddd.txt")
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		if lineNumber == lineCount {
			return fileScanner.Text()
		}
		lineCount++
	}
	defer file.Close()
	return ""
}
