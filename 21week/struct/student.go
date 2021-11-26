package main

import "fmt"

type student struct {
	name string
	age  int8
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{"小王子", 18},
		{"子怡", 19},
		{"baby", 20},
	}
	for _, stu := range stus {
		// m[stu.name] = &stu // 取stu的地址，但stu的值一直在变，为最后修改的值
		s := stu // student{name: stu.name, age: stu.age}
		m[stu.name] = &s
	}
	for k, v := range m {
		fmt.Printf("%v = %v\n", k, v)
	}
}
