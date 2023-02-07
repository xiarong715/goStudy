package main

import "fmt"

// 匿名函数：
// 匿名函数可以看作函数字面量
// 函数类型变量的地方可直接用匿名函数替代
// 匿名函数可直接赋值给函数类型变量
// 匿名函数可以作为实参
// 匿名函数可以作为返回值
// 匿名函数可以直接调用

// 匿名函数直接赋给函数类型变量
var sum = func(a int, b int) int {
	return a + b
}

// 函数类型作为形参
func do(f func(int, int) int, a, b int) int {
	return f(a, b)
}

// 匿名函数作为返回值
func wrap(op string) func(int, int) int {
	switch op {
	case "add":
		return func(a, b int) int {
			return a + b
		}
	case "sub":
		return func(a, b int) int {
			return a - b
		}
	default:
		return nil
	}
}

func main() {
	fmt.Printf("%d\n", do(sum, 1, 2)) //函数变量作为实参

	do(func(a, b int) int { // 匿名函数作为实参
		return a + b
	}, 1, 2)

	f := wrap("add") // 通过参数返回匿名函数
	fmt.Printf("%d\n", f(1, 2))

	// 匿名函数直接被调用
	func(a, b int) int {
		return a + b
	}(1, 2)

	fmt.Printf("hello world\n")
}
