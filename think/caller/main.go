package main

import (
	"caller/animal"
	"caller/cat"
	"caller/dog"
	"fmt"
)

func main() {
	a := animal.Animal{
		Name: "som",
		Age:  3,
	}
	c := cat.Cat{
		Animal: animal.Animal{
			Name: "linda",
			Age:  2,
		},
	}
	d := dog.Dog{
		Animal: animal.Animal{
			Name: "leo",
			Age:  2,
		},
	}

	a.SetName("dahuang")
	c.SetName("dahei")
	d.SetName("dacheng")

	// 指针调用者实现的接口，则要用指针类型调用接口的方法。值调用者实现的接口，则可用值类型也可用指针类型调用接口的方法。
	animal.Yell(&a)
	animal.Sleep(&a)

	animal.Yell(&c)
	animal.Sleep(&c)

	animal.Yell(&d)
	animal.Sleep(&d)

	fmt.Println("hello world.")
}
