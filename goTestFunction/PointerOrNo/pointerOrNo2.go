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

// note：若类型的方法会修改类型的值，那么接收者使用指针类型，否则会修改不成功
// 考虑到效率问题，使用值类型作为接收者，产生拷贝操作，所以建议所有方法的接收者类型都使用指针类型；
