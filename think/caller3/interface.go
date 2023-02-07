package main

import "fmt"

type Person struct {
	Name string
	Age  int8
}

type Student struct {
	Person
	ID int8
}

type Printer interface {
	Print()
}

func (p Person) getName() string {
	return p.Name
}

func (p *Person) Print() {
	fmt.Printf("name = %s, age = %d\n", p.Name, p.Age)
}

func (s *Student) Print() {
	fmt.Printf("name = %s, age = %d, id = %d\n", s.Name, s.Age, s.ID)
}

func Print(printer Printer) {
	printer.Print()
}
