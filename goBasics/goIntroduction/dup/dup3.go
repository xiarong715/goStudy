package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("dup3")
	fmt.Println("---------------------------")
	test()
	fmt.Println("---------------------------")
	test2()
}

func test() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3 test %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		fmt.Println(n, line)
	}
}

func test2() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3 test2 %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), " ") { // TODO
			counts[line]++
		}
	}
	for line, n := range counts {
		fmt.Println(n, line)
	}
}
