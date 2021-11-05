package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "how do you do"
	m := oneWordNum(sentence)
	for k, v := range m {
		fmt.Printf("%v = %v\n", k, v)
	}
}

func oneWordNum(str string) map[string]int {
	m := make(map[string]int, 10)
	s := strings.Split(str, " ")
	for _, v := range s {
		m[v] = m[v] + 1
	}
	return m
}
