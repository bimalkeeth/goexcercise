package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	file, err := ioutil.TempFile("", "hello_world_temp")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(file.Name())
	if _, err := file.Write([]byte("Hello world")); err != nil {
		fmt.Println(err)
	}
}
