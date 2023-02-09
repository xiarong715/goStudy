package main

import "fmt"

func recv(ch chan int) {
	fmt.Println("recv success", <-ch)
}

func recvCh(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("chan closed.")
			break
		}
		fmt.Println(v)
	}
}

// 转换为只能接收的单向通道
func recv2(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int) // 创建一个无缓冲通道
	go recv(ch)
	ch <- 5
	fmt.Println("send success")

	ch2 := make(chan int, 3) // 创建一个有3个缓冲的整型通道
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	close(ch2)
	recvCh(ch2)

	// for range 接收通道的值，当通道关闭后，接收完所有值自动退出
	ch3 := make(chan int, 2)
	ch3 <- 10
	ch3 <- 11
	close(ch3)
	for v := range ch3 {
		fmt.Println(v)
	}

	// 单向通道，只允许发送数据，或只允许接收数据
	// <- chan int 只允许发送数据的通道
	// chan <- int 只允许接收数据的通道
	// 在函数参数传递或赋值操作时，可将全向通道转换为单向通道，但不可反向转换。
	ch4 := make(chan int, 3)
	var ch5 chan<- int = ch4 // 转换为发送的单向通道
	ch5 <- 22
	ch5 <- 23
	close(ch4)
	recv2(ch4)

	// select 多路复用
	// 打印出10以内的奇数
	ch6 := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch6:
			fmt.Println(x)
		case ch6 <- i:
		}
	}
}

// 通道无缓冲时，发送方阻塞，直到接收方接收数据。反之，接收方阻塞，直到发送方发送数据。
// 无缓冲通道用于两goroutine的同步，通为同步通道。
// close关闭通道。一般是发送方关闭通道。关闭通道不是必须的。
// 通道关闭后，再向通道中发送数据会引发panic。
// 通道关闭后，可继续从通道中取出数据，直到通道中数据为空，取到的是通道类型的零值。
// 可通过取通道数据的第二个返回值判断通道是否被关闭。

// make 给 slice map chan 分配空间
