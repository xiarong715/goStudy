package main

import (
	"fmt"

	"github.com/go-macaron/inject"
)

type s1 interface{}
type s2 interface{}

func Print(name string, company s1, level s2, age int) {
	fmt.Printf("name=%v, company=%v, level=%v, age=%v\n", name, company, level, age)
}

func main() {
	Print("sim", "bytedance", "t3", 26)

	inj := inject.New()
	inj.Map("sim")
	inj.MapTo("bytedance", (*s1)(nil))
	inj.MapTo("t3", (*s2)(nil))
	inj.Map(26)

	inj.Invoke(Print)
}
