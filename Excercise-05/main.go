package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
	}
	//bs:=make([]byte,99999)
	//resp.Body.Read(bs)

	files, _ := os.Create("ddd.txt")
	defer files.Close()
	_, _ = io.Copy(files, resp.Body)

}
