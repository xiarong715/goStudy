package main

import (
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age  byte
	Say  string
}

// 此处的p 只是一个指针变量
func (p *Person) GetInfo() error {
	// p = &Person{}   // 改变的只是此处指针变量p的值
	p.Say = "say hello"

	return nil
}

func (p *Person) String() string {
	return p.Name + " " + fmt.Sprint(p.Age) + " " + p.Say
}

// 一切都是值传递，只是看这个值是普通值，还是地址值；

func main() {
	p := &Person{Name: "xiarong", Age: 20}
	if err := p.GetInfo(); err != nil {
		log.Fatal(err)
	}
	log.Println(p)
}
