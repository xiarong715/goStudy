package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", serve)
	http.ListenAndServe(":8080", nil)
}

func serve(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	fmt.Fprintf(rw, "\"%s\" is request", req.URL.Path)
}
