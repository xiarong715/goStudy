package main

import "log"

func main() {
	// 没有明确初始值的变量，会被赋予他们的零值
	var i int
	var b bool
	var s string
	var f float64
	// log.Println(i, b, s, f)
	log.Printf("%v %v %q %v", i, b, s, f)

	// 零值
	// 数值类型为 0
	// 布尔类型为 false
	// 字符串类型为 "" (空字符串)
}
