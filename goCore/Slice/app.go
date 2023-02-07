package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3, 4, 5}
	for _, v := range array {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")

	// 用数组创建切片
	s := array[1:3]
	fmt.Printf("%v\n", s)
	fmt.Printf("%d\n", len(s)) // 2
	fmt.Printf("%d\n", cap(s)) // 4,从当前指针开始算，到底层数组的最后一个元素

	fmt.Printf("-----------------\n")

	s = append(s, 6, 7, 8, 9)
	fmt.Printf("%v\n", s)
	fmt.Printf("%d\n", len(s))
	fmt.Printf("%d\n", cap(s))

	fmt.Printf("-----------------\n")

	// 用内置函数make创建切片
	slice := make([]int, 2, 5)
	slice[0] = 2
	slice[1] = 3
	slice = append(slice, 4) // 追加切片元素
	slice = append(slice, 5)
	slice = append(slice, 6)
	slice = append(slice, 7) // 追加元素，当容量不足时，会自动扩容
	fmt.Printf("%v\n", slice)
	fmt.Printf("%d\n", len(slice))
	fmt.Printf("%d\n", cap(slice))

	fmt.Printf("-----------------\n")

	sli := make([]int, 2, 3) // 当元素个数和容量相同时，可省略容量
	copy(sli, slice)         // copy 只会复制sli和slice中长度最小的
	fmt.Printf("%v\n", sli)
	fmt.Printf("%d\n", len(sli))
	fmt.Printf("%d\n", cap(sli))

	fmt.Printf("-----------------\n")

	fmt.Printf("hello world\n")
}
