package main // 特殊的包名，説明这是一个应用程序，而不是一个供其他包使用的库

import (
	"fmt" // 导入fmt包
	"os"
	"strings"
)

// Go语言和其他编程语言一样，一个大的程序是由很多小的基础构件组成的。
// 变量保存值，简单的加法和减法运算被组合成较复杂的表达式。
// 基础类型被聚合为数组或结构体等更复杂的数据结构。
// 然后使用if和for之类的控制语句来组织和控制表达式的执行流程。
// 然后多个语句被组织到一个个函数中，以便代码的隔离和复用。
// 函数以源文件和包的方式被组织。

func test1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func test2() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func test3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func test4() {
	fmt.Println(os.Args[1:])
}

// ./programstruct hello world

func main() { // 程序入口函数，整个程序的入口
	fmt.Println("hi, go program struct") // 调用fmt包中的函数
	fmt.Println("------------------------------")
	test1()
	fmt.Println("------------------------------")
	test2()
	fmt.Println("------------------------------")
	test3()
	fmt.Println("------------------------------")
	test4()
	fmt.Println("------------------------------")
}
