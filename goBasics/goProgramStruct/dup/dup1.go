package main

import (
	"bufio"
	"fmt"
	"os"
)

// 从屏幕输入多行字符串，统计重复出现的行，打印出现的次数及其行内容

func test() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() && input.Text() != "-1" {
		counts[input.Text()]++
	}
	fmt.Println("-----------------------------")
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func main() {
	fmt.Println("input multiline string: ")
	test()
}
