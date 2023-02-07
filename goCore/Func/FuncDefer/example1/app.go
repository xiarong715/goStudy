package main

import "fmt"

func main() {
	// 一个函数内可注册多个延迟调用
	// 延迟调用按照【先进后出】的顺序在函数返回前被调用
	defer func() {
		fmt.Printf("first defer\n")
	}()

	defer func() {
		fmt.Printf("second defer\n")
	}()

	defer func() {
		fmt.Printf("three defer\n")
	}()

	fmt.Printf("hello world\n")
}
