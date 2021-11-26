package main

import (
	"fmt"
)

type Animal struct {
	name string
}

func (a Animal) Eat() {
	fmt.Printf("%s is eating\n", a.name)
}

type Cat struct {
	Animal
}

func main() {
	cat := Cat{
		Animal: Animal{
			name: "cat",
		},
	}
	cat.Eat()
}
