package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	*Person
	Number int
}

func main() {
	per := Person{
		name: "xia",
		age:  18,
	}
	per.age = 19

	stu := Student{
		Person: &per,
		Number: 1,
	}

	fmt.Printf("%v\n", stu)

	fmt.Printf("hello world\n")
}
