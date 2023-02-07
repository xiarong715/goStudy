package main

import "fmt"

type Printer interface {
	Print()
}

type Person struct {
	name string
	age  int8
}

func (p Person) getName() string {
	return p.name
}

func (p Person) getAge() int8 {
	return p.age
}

func (p Person) Print() {
	fmt.Printf("name = %s, age = %d\n", p.name, p.age)
}
