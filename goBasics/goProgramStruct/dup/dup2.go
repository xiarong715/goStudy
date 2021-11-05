package main

import (
	"bufio"
	"fmt"
	"os"
)

func test() {
	counts := make(map[string]int)
	files := os.Args[1:]
	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(file, counts)
		file.Close()
	}
	for line, n := range counts {
		fmt.Println(n, line)
	}
}

func countLines(file *os.File, counts map[string]int) { // map 是引用传递
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func main() {
	fmt.Println("dup2")
	test()
}
