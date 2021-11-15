package main

import (
	"bufio"
	"fmt"
	"os"
)

// ./dup2_exercise TESTFILE TESTFILE2

// 统计打印文件中重复的行，并打印所在的文件
func main() {
	fmt.Println("dup2_exercise")
	test()
}

func test() {
	counts := make(map[string](map[string]int))
	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2_exercise %v\n", err)
			continue
		}
		countLines(file, filename, counts)
		file.Close()
	}
	for filename, m := range counts {
		fmt.Println("------------------")
		fmt.Println(filename)
		for line, n := range m {
			fmt.Println(n, line)
		}
	}
}

func countLines(file *os.File, filename string, count map[string](map[string]int)) {
	input := bufio.NewScanner(file)
	if count[filename] == nil {
		count[filename] = make(map[string]int)
	}
	for input.Scan() {
		count[filename][input.Text()]++
	}
}
