package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3, 4, 5} // 创建数组后，长度固定，可修改元素值
	array[0] = 2

	for _, v := range array { // 遍历数组
		fmt.Printf("%d ", v)
	}
	fmt.Print("\n")

	fmt.Printf("%d\n", array[0])
	fmt.Printf("%d\n", len(array))
}
