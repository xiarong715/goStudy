package main

import (
	"bufio"
	"fmt"
)

type WordCounter int
type LineCounter int
type scanFunc func([]byte, bool) (int, []byte, error)

func retCount(p []byte, fn scanFunc) (int, error) {
	count := 0
	for {
		advance, token, _ := fn(p, true)
		if len(token) == 0 {
			break
		}
		p = p[advance:]
		count++
	}
	return count, nil
}

func (c *WordCounter) Write(p []byte) (int, error) {
	cnt, _ := retCount(p, bufio.ScanWords)
	*c += WordCounter(cnt)
	return cnt, nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	cnt, _ := retCount(p, bufio.ScanLines)
	*c += LineCounter(cnt)
	return cnt, nil
}

func main() {
	fmt.Println("word counter")
	m, n, _ := bufio.ScanWords([]byte("hello beautiful world"), false)
	fmt.Println(m, string(n))
	var wc WordCounter
	fmt.Fprintf(&wc, "%s", "hello beautiful world")
	fmt.Println(wc)
	fmt.Println("------------------------")
	var lc LineCounter
	fmt.Fprintf(&lc, "%s", "hello \nbeautiful \nworld")
	fmt.Println(lc)
}
