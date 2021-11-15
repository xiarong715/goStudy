package main

import (
	"fmt"

	"goMethod/geometry"
)

func main() {
	fmt.Println("geometry")
	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{X: 3, Y: 4}
	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))
}
