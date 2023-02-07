package main

import "fmt"

var a = 0

func fa() func(int) int {
	return func(i int) int {
		fmt.Println(&a, a)
		a = a + i
		return a
	}
}

func main() {
	f := fa()
	g := fa()
	fmt.Println(f(1))
	fmt.Println(f(1))
	fmt.Println(f(1))
	fmt.Println(g(1))

	fmt.Println("hello world")
}
