package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type WordCounter int
type LineCounter int
type ByteCounter int

func retCount(p []byte, fn bufio.SplitFunc) (int, error) {
	cnt := 0
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(fn)
	for scanner.Scan() {
		cnt++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return 0, err
	}
	return cnt, nil
}

func (c *WordCounter) Write(p []byte) (int, error) {
	cnt, err := retCount(p, bufio.ScanWords)
	if err != nil {
		return cnt, err
	}
	*c += WordCounter(cnt)
	return cnt, nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	cnt, err := retCount(p, bufio.ScanLines)
	if err != nil {
		return cnt, err
	}
	*c += LineCounter(cnt)
	return cnt, nil
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	cnt, err := retCount(p, bufio.ScanBytes)
	if err != nil {
		return cnt, err
	}
	*c += ByteCounter(cnt)
	return cnt, nil
}

func main() {
	fmt.Println("counter bytes, lines")
	var wc WordCounter
	var lc LineCounter
	var bc ByteCounter
	content := "hello world.\n i comming"
	fmt.Fprintf(&wc, "%v", content)
	fmt.Println(wc)
	fmt.Println("------------------------")
	fmt.Fprintf(&lc, "%v", content)
	fmt.Println(lc)
	fmt.Println("------------------------")
	fmt.Fprintf(&bc, "%v", content)
	fmt.Println(bc)
}
