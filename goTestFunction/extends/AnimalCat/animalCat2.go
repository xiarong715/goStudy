package main

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) Eat() {
	fmt.Printf("%s is eating\n", a.Name)
}

type Cat struct {
	*Animal
}

func main() {
	cat := &Cat{
		Animal: &Animal{
			Name: "cat",
		},
	}
	cat.Eat()
}
