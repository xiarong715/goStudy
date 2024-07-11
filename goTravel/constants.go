package main

import (
	"log"
	"testing"
)

// 常量可以是：字符、字符串、布尔值和数值；
// 不可以用 := 语法声明常量

const PI = 3.14

func TestConstants(t *testing.T) {
	const World = "世界"
	const Truth = true
	log.Println("Hello", World)
	log.Println("PI =", PI)
	log.Println("Truth =", Truth)

}
