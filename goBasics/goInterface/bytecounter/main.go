package main

import (
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	fmt.Println("byte counter")
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0 // reset counter
	fmt.Fprintf(&c, "hello, %s", "world")
	fmt.Println(c)
}
