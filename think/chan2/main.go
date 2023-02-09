package main

import (
	"fmt"
	"sync"
	"time"
)

func demo1() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	wg.Add(1)
	go func() {
		for {
			v, ok := <-ch
			if !ok { // 必须要判断通道是否关闭，否则goroutine不能退出
				break
			}
			fmt.Println(v)
		}
		wg.Done()
	}()
	wg.Wait()
}

func demo2() {
	ch := make(chan string)
	go func() {
		time.Sleep(3 * time.Second) // 耗时更久
		ch <- "job result"          // 发送被阻塞
	}()

	select {
	case x := <-ch:
		fmt.Println(x)
	case t := <-time.After(time.Second): // 1s后从通道中读取当前的时间。耗时短
		fmt.Println(t)
		return // goroutine发送阻塞，没来得及从通道接收数据就退出了，导致goroutine泄露
	}
}

func main() {
	demo1()
	demo2()
	fmt.Println("hello chan")
}
