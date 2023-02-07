package main

type Student struct {
	Person
	ID int8
}

func (s Student) getID() int8 {
	return s.ID
}

func (s *Student) setID(id int8) {
	s.ID = id
}
