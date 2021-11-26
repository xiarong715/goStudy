package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

func (p *Point) ScaleBy(factor int) {
	p.X *= float64(factor)
	p.Y *= float64(factor)
}

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// type ColoredPoint struct { // "has a" Point
// 	Point  // 匿名字段
// 	Color color.RGBA
// }

// func main() {
// 	var cp ColoredPoint
// 	cp.X = 1
// 	cp.Point.Y = 2
// 	fmt.Println(cp.X, cp.Y)

// 	red := color.RGBA{255, 0, 0, 255}
// 	blue := color.RGBA{0, 0, 255, 255}
// 	p := ColoredPoint{Point{1, 1}, red}
// 	q := ColoredPoint{Point{5, 4}, blue}
// 	fmt.Println(p.Point.Distance(q.Point))
// 	p.ScaleBy(2)
// 	q.ScaleBy(2)
// 	fmt.Println(p.Distance(q.Point))
// }

// type ColoredPoint struct {
// 	*Point // 匿名字段
// 	Color  color.RGBA
// }

// func main() {
// 	red := color.RGBA{255, 0, 0, 255}
// 	blue := color.RGBA{0, 0, 255, 255}
// 	p := ColoredPoint{&Point{1, 1}, red}
// 	q := ColoredPoint{&Point{5, 4}, blue}
// 	fmt.Println(p.Point.Distance(*q.Point))
// 	p.ScaleBy(2)
// 	q.ScaleBy(2)
// 	fmt.Println(p.Distance(*q.Point))
// }

type ColoredPoint struct {
	Point
	color.RGBA
}

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	p := ColoredPoint{Point{1, 1}, red}
	q := ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Point, p.RGBA)
	fmt.Println(q.Point, q.RGBA)
	p.ScaleBy(2)
	p.Distance(q.Point)
	p.RGBA.RGBA() // 冲突时，需指明字段
}
