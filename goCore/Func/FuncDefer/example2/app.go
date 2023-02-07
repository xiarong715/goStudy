package main

import "fmt"

// defer函数注册后才能被执行
func main() {
	fmt.Printf("hello world\n")

	return

	// 在return后的代码得不到执行，也就没有注册，所以defer函数得不到调用
	defer func() {
		fmt.Printf("defer\n")
	}()
}
