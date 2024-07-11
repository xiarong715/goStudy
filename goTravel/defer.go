package main

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	// defer语句会将函数推迟到外层函数返回之后执行；
	// 推迟调用的函数参数会立即求值，直到外层函数返回前该函数都不会被调用
	defer fmt.Println("World")

	fmt.Println("Hello")
}
