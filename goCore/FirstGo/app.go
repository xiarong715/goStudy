package main // 每个源文件都有一个包名，且都在源文件的开始声明包。main是可执行程序的包名

import (
	"fmt" // 导入 fmt 包，用输入输出操作
	"unsafe"
)

// main函数是执行程序的入口函数
func main() {
	fmt.Printf("hello world\n") // 调用 fmt中的 Printf函数，输出文字
	var a int = 100
	fmt.Printf("--------\n")
	fmt.Printf("%d\n", unsafe.Sizeof(a))

	var b int = 5
	p := &b
	fmt.Print(*p)
}

// 程序员水平高的表现
// 思考问题全面、细致
// 足够的好奇心，更高的要求
// 对结果充分负责

// 做技术感到快乐的点
// 上线的时候，
// 达到预期

// 偶像
//

// 技术影响力是什么
