package main

import (
	"log"
	"testing"
)

// 变量声明可以赋初始值，每个变量对应一个；
// 变量初始值存在，可以省略类型；变量会从初始值中获取类型；
var cs, javas, pythons = true, false, false

func TestVariablesWithInit(t *testing.T) {
	var i, j = 0, 1
	log.Println(i, j, cs, javas, pythons)
}
