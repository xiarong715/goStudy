package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sayHi() {
	fmt.Println("hi.")
	wg.Done() // 减少一个等待的goroutine的计数
}

func hello(i int) {
	defer wg.Done()
	fmt.Printf("%d\n", i)
}

func main() {
	wg.Add(1)  // 增加一个等待的goroutine的计数
	go sayHi() // 启动一个goroutine
	fmt.Println("hello world.")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}

	// 闭包的坑
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("t", i) // 产生闭包。没通过参数传值，执行时才确定值。可能得不到想要的结果。
		}()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("s", i) // 好的处理方式：通过参数传入值，不使用闭包
		}(i)
	}

	wg.Wait() // 等待goroutine的计数为0

	// 环境在变化，且闭包函数不能马上被执行时，用参数接收环境值。
	// 最好的做法是：一直用参数接收传入的值。
}
