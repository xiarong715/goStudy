package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

type Ops func(int, int) int

func do(f Ops, a, b int) int { // 插座
	return f(a, b) // 函数类型变量可以直接用来进行函数调用
}

func main() {
	fmt.Printf("%d\n", do(add, 2, 3)) // 函数名add作为do函数的参数，函数名就是该函数类型变量名
	fmt.Printf("%d\n", do(sub, 2, 3))

	fmt.Printf("%d\n", add(1, 2)) // 直接通过函数调用函数
	f := add                      // 函数名赋值给函数类型变名
	fmt.Printf("%d\n", f(1, 2))   // 通过函数类型变量调用函数

	fmt.Printf("hello world\n")
}
