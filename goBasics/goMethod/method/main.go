package main

import (
	"fmt"

	"method/geometry"
)

func main() {
	fmt.Println("geometry")
	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{X: 3, Y: 4}
	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))
	fmt.Println("--------------------------------")
	prim := geometry.Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Println(geometry.PathDistance(prim))
	fmt.Println(prim.Distance())
}
