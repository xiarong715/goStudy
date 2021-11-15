package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server1")
	test()
}

func test() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8888", nil)
}

func handle(w http.ResponseWriter, request *http.Request) {
	// fmt.Println(request.URL.Path)
	fmt.Fprintf(w, "the request url: %s", request.URL.Path)
}
