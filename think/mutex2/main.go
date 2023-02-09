package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x       int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
	rwMutex sync.RWMutex
)

func writeWithLock() {
	mutex.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	mutex.Unlock()
	wg.Done()
}

func readWithLock() {
	mutex.Lock()
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()
}

func writeWithRWLock() {
	rwMutex.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	rwMutex.Unlock()
	wg.Done()
}

func readWithRWLock() {
	rwMutex.RLock()
	time.Sleep(time.Millisecond)
	rwMutex.RUnlock()
	wg.Done()
}

func do(wf, rf func(), wc, rc int) {
	start := time.Now()
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go wf()
	}
	for i := 0; i < rc; i++ {
		wg.Add(1)
		go rf()
	}
	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("x:%v, cost:%v\n", x, cost)
}

func main() {
	// 互斥锁和读写锁用在读多写少的场景，测试性能。
	do(writeWithLock, readWithLock, 10, 1000)
	do(writeWithRWLock, readWithRWLock, 10, 1000)
	fmt.Println("hello mutex2")
}
