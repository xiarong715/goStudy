package main

import (
	"fmt"
	"os"
)

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

// 重写Animal的Eat方法
func (c *Cat) Eat() {
	fmt.Printf("children %s is eating\n", c.Name)
}

func (c *Cat) String() string {
	return c.Name
}

func newCat(name string) *Cat {
	return &Cat{
		Animal: newAnimal(name),
	}
}

func main() {
	cat := newCat("cat")
	cat.Eat()
	fmt.Fprintln(os.Stdout, cat)
}
