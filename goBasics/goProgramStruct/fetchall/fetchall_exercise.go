package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("fetchall_exercise")
	test()
}

func test() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		for i := 0; i < 2; i++ {
			go fetch(url, ch)
		}
	}
	for range os.Args[1:] {
		for i := 0; i < 2; i++ {
			fmt.Println(<-ch)
		}
	}
	fmt.Printf("time elasped %.2fs\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("fetch get %v\n", err)
	}
	numBytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("fetch read %v\n", err)
	}
	ch <- fmt.Sprintf("%.2fs\t%d\t%s", time.Since(start).Seconds(), numBytes, url)
}
