package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// go build fetch.go
// ./fetch https://baidu.com

func main() {
	fmt.Println("fetch")
	test()
}

func test() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v\n", err)
			os.Exit(1)
		}
		data, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "read %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", data)
	}
}
