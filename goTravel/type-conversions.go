package main

import "log"

func main() {
	// 与C不同，go在不同类型之间赋值时需显示转换
	var x, y int = 5, 6
	var ux uint = uint(x)
	var uy uint = uint(y)

	log.Println(x, y, ux, uy)
}
