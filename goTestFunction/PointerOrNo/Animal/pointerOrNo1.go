package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) setName(name string) {
	a.Name = name
}

func (a Animal) String() string {
	return a.Name
}

type Cat struct {
	Animal
}

func main() {
	cat := Cat{
		Animal: Animal{
			Name: "cat",
		},
	}
	fmt.Println(cat)
	cat.setName("maozi") // 因为接收者为值类型，所以修改字段失败
	fmt.Println(cat)
}
