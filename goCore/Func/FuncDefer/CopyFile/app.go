package main

import "fmt"

func swap(a, b int) (x, y int) {
	x = b
	y = a
	return
}

func main() {
	x, y := swap(2, 3)
	fmt.Printf("%d %d\n", x, y)
	fmt.Printf("hello world\n")
}
