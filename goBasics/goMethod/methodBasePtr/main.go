package main

import (
	"fmt"
	"methodBasePtr/geometry"
)

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{3, 4}
	fmt.Println((p).Distance(q))

	fmt.Println("-------------------")

	prim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println((prim).Distance())
	fmt.Println(prim.Sss())
}
