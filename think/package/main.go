package main

import (
	"pack/interf"
	"pack/person"
	"pack/student"
)

func Print(printer interf.Printer) {
	printer.Print()
}

func main() {
	per := person.Person{
		Name: "heo",
		Age:  18,
	}
	stu := student.Student{
		Person: person.Person{
			Name: "tom",
			Age:  20,
		},
		ID: 1,
	}
	Print(per)
	Print(stu)
}
