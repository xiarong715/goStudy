package main

import "fmt"

type IAnimal interface {
	Eat()
}

type Animal struct {
	Name string
}

func (a *Animal) Eat() {
	fmt.Printf("%s is eating\n", a.Name)
}

func newAnimal(name string) *Animal {
	return &Animal{
		Name: name,
	}
}

type Cat struct {
	*Animal
}

func newCat(name string) *Cat {
	return &Cat{
		Animal: newAnimal(name),
	}
}

func main() {
	cat := newCat("cat")
	cat.Eat()
}
