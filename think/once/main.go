package main

import (
	"fmt"
	"once/singleton"
)

func main() {
	inst1 := singleton.GetInstance()
	inst2 := singleton.GetInstance()
	if inst1 == inst2 {
		fmt.Println("inst1 == inst2")
	}
	fmt.Println("hello once")
}
