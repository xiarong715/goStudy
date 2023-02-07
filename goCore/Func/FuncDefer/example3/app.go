package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Printf("defer\n")
	}()

	fmt.Printf("hello world\n")

	os.Exit(0) // 手动调用os.Exit(int)退出进程，defer函数将不被调用
}
