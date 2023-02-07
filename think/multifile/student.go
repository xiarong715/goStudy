package main

import (
	"fmt"
)

type Student struct {
	Person
	id int8
}

func (s Student) getID() int8 {
	return s.id
}

func (s Student) Print() {
	fmt.Printf("name = %s, age = %d, id = %d\n", s.Person.getName(), s.Person.getAge(), s.getID())
}
