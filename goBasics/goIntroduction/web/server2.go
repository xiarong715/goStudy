package main

import (
	"fmt"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int64

func main() {
	http.HandleFunc("/", Handle)
	http.HandleFunc("/count", Count)
	http.ListenAndServe(":8888", nil)
}

func Handle(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "the request path: %s", r.URL.Path)
}

func Count(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "the Count: %d", count)
	mu.Unlock()
}
