package main

import (
	"fmt"
	"log"
)

// Operation operation
type Operation interface {
	count(int, int)
}

// Add Sub 实现了Operation接口的方法，因此Add Sub具有Operation的特性

// Add add
type Add struct {
}

func (*Add) count(a, b int) {
	fmt.Printf("a+b=%d\n", a+b)
}

// Sub sub
type Sub struct {
}

func (*Sub) count(a, b int) {
	fmt.Printf("a-b=%d\n", a-b)
}

// Context context
// 嵌入Operation类型，要嵌入一个类型，只需要声明这个类型的名字
// 外部类型Operation成为外部类型Context的内部类型
// 内部类型相关的标识符（字段、方法、接口等）会提升到外部类型上
// 外部类型声明了的，就不会被提升
// 则Context结构拥有了Operation接口类型的所有方法
type Context struct {
	Operation
}

// NewContext new Context
// 使用接口作为参数，产生多态的效果
func NewContext(operation Operation) *Context {
	return &Context{operation}
}

func main() {
	log.Printf("strategy pattern")
	state := 2
	a := 3
	b := 3
	oper := make(map[int]Operation)
	oper[1] = &Add{}
	oper[2] = &Sub{}

	c := NewContext(oper[state])
	c.count(a, b)
}
