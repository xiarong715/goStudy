package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 针对一些有导出字段的结构体，初始化（创建结构体变量）时使用{}的形式
	b := new(bytes.Buffer) //bytes.Buffer{}  bytes.Buffer结构体中的字段一个也没导出，一般使用new的方式创建结构体变量
	b.Write([]byte("hello world"))
	fmt.Println(string(b.Bytes()))
}
