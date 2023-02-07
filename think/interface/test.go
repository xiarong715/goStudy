package main

import "fmt"

type Printer interface {
	Print()
}

type Person struct {
	name string
	age  int8
}

func (p Person) print() {
	fmt.Printf("name = %s, age = %d\n", p.name, p.age)
}

func (p Person) getName() string {
	return p.name
}

func (p Person) getAge() int8 {
	return p.age
}

type Student struct {
	Person // 匿名定义
	id     int8
}

func (s Student) getID() int8 {
	return s.id
}

func (s Student) Print() {
	fmt.Printf("name = %s, age = %d, id = %d\n", s.getName(), s.getAge(), s.id) // 匿名调用，更有继承的感觉
}

// 实现Printer接口的Print函数，实现接口中所有函数才算实现了接口。
// Printer接口作为参数时，实现了Printer接口的结构体都可以作为参数。实现了Printer接口的结构体有了Printer的特性。
// 接口中包含一系列的方法，空接口 interface{} 中没有方法。go所有数据类型都实现了 interface{}。
// 实现接口是实现多态的前提。当然在golang中不叫多态。
func (p Person) Print() {
	fmt.Printf("name = %s, age = %d\n", p.name, p.age)
}

func Print(printer Printer) { // 实现多态
	printer.Print()
}

// 类型断言
func assertType(x interface{}) {
	_, ok := x.(Person)
	if ok {
		fmt.Printf("x is Person type.\n")
	} else {
		fmt.Printf("x is not Person type, is %T type\n", x)
	}
}

func main() {
	per := Person{
		name: "tom",
		age:  18,
	}
	per.print()

	per.age = 20 // 在包内可访问与修改小写字母开头的成员
	per.name = "switch"
	fmt.Printf("name = %s, age = %d\n", per.getName(), per.getAge())

	per.age = 22
	per.name = "hansom"
	per.Print()

	stu := Student{
		Person: Person{
			name: "jim",
			age:  25,
		},
		id: 1,
	}
	stu.Print()

	stu.name = "jom" // 更有继承的感觉
	stu.age = 20
	Print(per)
	Print(stu)

	assertType(per)
	assertType(stu)

	fmt.Printf("%s\n", "Hello world.")
}
