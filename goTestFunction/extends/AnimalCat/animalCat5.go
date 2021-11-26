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

// 重写Animal的Eat方法
func (c *Cat) Eat() {
	fmt.Printf("children %s is eating\n", c.Name)
}

func newCat(name string) *Cat {
	return &Cat{
		Animal: newAnimal(name),
	}
}

type Dog struct {
	*Animal
}

// 重写Animal的Eat方法
func (d *Dog) Eat() {
	fmt.Printf("children %s is eating\n", d.Name)
}

func newDog(name string) *Dog {
	return &Dog{
		Animal: newAnimal(name),
	}
}

// 参数多态
func check(animal IAnimal) {
	animal.Eat()
}

func main() {
	animal := newAnimal("animal")
	cat := newCat("maozi")
	dog := newDog("gouzi")
	check(animal)
	check(cat)
	check(dog)
}
