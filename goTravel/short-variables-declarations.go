package main

import (
	"log"
)

// 函数外的语句都必须以关键字开始(var、func等等)，因此 := 不在能在函数外使用；

func main() {

	// 在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明
	var i, j = 5, 6
	res := "str"
	c, java, python := true, false, "no!"
	log.Println(i, j, res, c, java, python)
}
