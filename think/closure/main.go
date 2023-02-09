package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		func() {
			fmt.Println(i)
		}()
	}
	fmt.Println("hello closure")
} // 待完成
