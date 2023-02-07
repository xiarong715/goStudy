package main

import "fmt"

func main() {
	// 定义数组
	x := [3]int{1, 2, 3}
	y := [5]int{1, 2, 3, 4, 5}
	fmt.Println(x, y)

	// 定义切片
	s := y[0:2]
	fmt.Printf("%T\n", s)
	fmt.Println(s)
	fmt.Printf("len = %d, capacity = %d\n", len(s), cap(s)) // 2, 5

	s = x[2:3]
	fmt.Printf("%T\n", s)
	fmt.Println(s)
	fmt.Printf("len = %d, capacity = %d\n", len(s), cap(s)) // 1, 1

	// 切片的扩容策略：1，2，4，8，16，……   即2的指数次，2^n，n为0，1，2，3，4，……
	// 切片扩容的本质：更换更大容量的底层数组。
	s = append(s, 4)
	fmt.Printf("len = %d, capacity = %d\n", len(s), cap(s)) // 2, 2

	// 切片是引用类型，当两个切片引用同一个底层数组时，一个切片值的修改会影响另一个切片的值。
	s1 := []int{1, 2, 3, 4}
	s2 := s1
	fmt.Println(s1, s2)
	s2[0] = 2000
	fmt.Println(s1, s2)

	// 通过拷贝切片，消除上面的问题。
	s1 = []int{1, 2, 3, 4}
	s2 = make([]int, 4)
	copy(s2, s1)
	s2[0] = 2000
	fmt.Println(s1, s2)

	// append向切片中添加元素，可一次添加一个元素，也可一次添加多个元素，还可添加一个切片
	s1 = append(s1, 5)       // 添加一个元素
	s1 = append(s1, 6, 7, 8) // 添加多个元素
	s1 = append(s1, s2...)   // 添加一个切片
	fmt.Println(s1)

	// 切片的遍历
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%d ", s1[i])
	}
	fmt.Println()

	for _, v := range s1 {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// 切片的比较
	// 一个nil的切片没有底层数组。
	// 判断一个切片没有元素：len(s) == 0
	// 判断一个切片没有底层数组：s == nil
	// 切片之间不能比较。

	// 删除切片的元素
	// 删除切片中索引为index的元素：s = append(s[:index], s[index+1:]...)
	// 如删除s1中索引为8的元素
	s1 = append(s1[:8], s1[9:]...)
	fmt.Println(s1)

	// 切片的本质
	// 切片类型有三个信息：切片指向的底层数组的指针，切片的长度，切片的容量。
}
