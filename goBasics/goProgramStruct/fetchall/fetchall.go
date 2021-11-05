package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("fetchall")
	test()
}

// 当一个goroutine尝试在一个channel上send或receice消息时，这个goroutine会阻塞在调用处；
// 直到另一个goroutine往这个channel中receive或send消息，这样两个goroutine才会继续执行channel完成之后的逻辑；

func test() {
	start := time.Now()
	ch := make(chan string) // create and assign channel
	for _, url := range os.Args[1:] {
		go getUrl(url, ch)
	}
	for range os.Args[1:] { // TODO
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("time elasped %.2fs\n", time.Since(start).Seconds())
}

func getUrl(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("getUrl get %v\n", err) // send to channel ch
		return
	}
	numBytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("getUrl read %v\n", err) // send to channel ch
		return
	}

	ch <- fmt.Sprintf("%.2fs\t%d\t%s", time.Since(start).Seconds(), numBytes, url) // send to channelch
}
