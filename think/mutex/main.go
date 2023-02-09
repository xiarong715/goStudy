package main

import (
	"fmt"
	"sync"
)

var (
	x   int64
	wg  sync.WaitGroup
	m   sync.Mutex   // 互斥锁
	rwm sync.RWMutex // 读写锁
)

func add() {
	for i := 0; i < 50000; i++ {
		x = x + 1 // 多个goroutine抢资源，没有安全措施，数据会错乱
	}
	wg.Done()
}

func add2() {
	for i := 0; i < 50000; i++ {
		m.Lock() // 在临界区前加锁
		x = x + 1
		m.Unlock() // 在临界区访问后释放锁
	}
	wg.Done()
}

func main() {
	wg.Add(4)
	// go add()
	// go add()
	// go add()
	// go add()
	go add2()
	go add2()
	go add2()
	go add2()
	wg.Wait()
	fmt.Println(x)
}
