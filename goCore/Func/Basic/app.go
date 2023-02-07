package main

import "fmt"

// 多个相邻的参数类型相同，可简写
// 命名返回参数变量sum，也是函数的局部变量
func add(a, b int) (sum int) {
	sum = a + b
	return sum // 当给命名返回值变量赋值后，return可直接返回，不用带参数
}

func add1(a, b int) (sum int) {
	anonymous := func(x, y int) int { // 函数中支持嵌套匿名函数
		return x + b
	}
	return anonymous(a, b)
}

func swap(a, b int) (int, int) { // 两个匿名返回参数,当有我个返回值时，返回列表的括号不能省略
	return b, a
}

func chvalue(a int) int {
	a = a + 1
	return a
}

func chpointer(a *int) {
	*a = *a + 1
}

func main() {
	x := 2
	y := 3
	fmt.Printf("%d\n", add(x, y))
	fmt.Printf("%d\n", add1(x, y))
	a, b := swap(x, y)
	fmt.Printf("%d %d\n", a, b)

	c := 10
	chvalue(c) // 值传递
	fmt.Printf("%d\n", c)

	chpointer(&c) // 值传递，此处的值是地址
	fmt.Printf("%d\n", c)

	fmt.Printf("hello world\n")
}
