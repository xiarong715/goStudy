package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int)
var mutex = sync.Mutex{}

var syncmap = sync.Map{} // sync包中提供的Map，是goroutine安全的

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func demo1() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			mutex.Lock() // 不加锁时，多个goroutine对map会产生同时写的情况，产生错误
			set(key, n)
			mutex.Unlock()
			fmt.Printf("key = %v, value = %v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func demo2() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			syncmap.Store(key, n)
			value, _ := syncmap.Load(key)
			fmt.Printf("key = %v, value = %v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {
	demo1()
	fmt.Println("-------------------")
	demo2()
	fmt.Println("hello sync map")
}
