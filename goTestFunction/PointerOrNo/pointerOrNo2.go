package main

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) setName(name string) {
	a.Name = name
}

func (a *Animal) String() string {
	return a.Name
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
	fmt.Println(cat)
	cat.setName("maozi") // 接收者为指针类型，修改成功
	fmt.Println(cat)
}
