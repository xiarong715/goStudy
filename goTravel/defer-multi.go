package main

import (
	"fmt"
	"testing"
)

func TestDeferMulti(t *testing.T) {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
	fmt.Print(1, 2)
}
