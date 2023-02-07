package student

import (
	"fmt"
	"pack/person"
)

type Student struct {
	person.Person
	ID int8
}

func (s Student) getID() int8 {
	return s.ID
}

func (s Student) Print() {
	fmt.Printf("name = %s, age = %d, id = %d\n", s.Person.GetName(), s.Person.GetAge(), s.getID())
}
