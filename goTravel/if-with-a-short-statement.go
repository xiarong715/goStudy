package main

import (
	"log"
	"math"
)

func main() {
	log.Println(
		pow(2, 3, 10),
		pow(3, 3, 20),
	)
}

func pow(x, n, lim float64) float64 {
	// if 语句，在条件语句前执行一个语句；
	// if 语句中定义的变量，只能在if结构中使用；
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
