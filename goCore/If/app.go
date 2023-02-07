package main

import "fmt"

func main() {
	x := 2
	y := 3
	if x = 4; x <= y {
		fmt.Printf("x < y\n")
	} else if x > y {
		fmt.Printf("x > y\n")
	} else {
		fmt.Printf("x == y\n")
	}

	fmt.Printf("hello world\n")
}
