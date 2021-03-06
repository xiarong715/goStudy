package main

import (
	"fmt"
	"sync"
)

// var (
// 	mu      sync.Mutex
// 	mapping = make(map[string]string)
// )

// func LookUp(key string) string {
// 	mu.Lock()
// 	v := mapping[key]
// 	mu.Unlock()
// 	return v
// }

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func LookUp(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func main() {
	fmt.Println(LookUp("lsy"))
}
