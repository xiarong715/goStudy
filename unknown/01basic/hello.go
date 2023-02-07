package main

import (
	"fmt"
	. "fmt"  // 省略调用，不建议使用
	io "fmt" // 导入成 io
)

// 自定义类型
type myInt int
type gostruct struct{}
type gointerface interface{}

// 类型别名
type intAlias = int

// 大小写决定常量、变量、函数、接口、类型、结构是否可被外部使用

// main包的main函数：main.main
// main包中必须要有main函数
// 只有main包才能编译成执行程序
func main() {
	io.Println("hello world")
	var i myInt = 3
	fmt.Println(i)
	Println(i) // 省略调用

	var a myInt
	var b intAlias
	fmt.Printf("%T\n", a) // 新的数据类型 main.myInt
	fmt.Printf("%T\n", b) // 类型还是 int
}
