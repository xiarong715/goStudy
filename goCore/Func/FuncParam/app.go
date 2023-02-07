package main

import "fmt"

func sum(param ...int) (sum int) { // 命名的返回参数，当为命名返回参数时，（）也不能省略
	// 不定参数，在函数内相当于切片，对切片的所有操作都可对不定参数进行操作
	fmt.Printf("len = %d\n", len(param))
	fmt.Printf("cap = %d\n", cap(param))
	for _, v := range param {
		sum += v
	}
	return // 当返回参数有名时，给返回参数赋值后，可直接return返回，可不带参数
}

func main() {
	fmt.Printf("%d\n", sum(1, 2, 3, 4, 5, 6))

	slice := []int{2, 3, 4}
	fmt.Printf("%d\n", sum(slice...))

	fmt.Printf("hello world")
}
