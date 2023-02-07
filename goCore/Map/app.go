package main

import "fmt"

func main() {
	// 创建map
	// 用字面量创建map
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for i, v := range m {
		fmt.Printf("%s: %d\n", i, v)
	}
	fmt.Printf("len = %d\n", len(m))

	m2 := make(map[int]string, 2)
	for i, v := range m2 {
		fmt.Printf("%d: %s\n", i, v)
	}
	m2[1] = "slice"
	m2[2] = "map"
	m2[3] = "array"
	fmt.Printf("len = %d\n", len(m2))

	type User struct {
		name string
		age  int
	}

	mp := make(map[int]User)
	anna := User{
		name: "anna",
		age:  19,
	}
	mp[0] = anna
	// mp[0].age = 20  // error    不能通过map引用直接修改值中的某个元素
	anna.age = 20 // 单独修改值
	mp[0] = anna  // 整体替代

	fmt.Printf("hello world\n")
}
