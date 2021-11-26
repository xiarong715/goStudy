package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor int) {
	p.X *= float64(factor)
	p.Y *= float64(factor)
}

func main() {
	// 方法“值”
	p := Point{1, 1}
	q := Point{5, 4}
	distanceFromP := p.Distance // 选择器取到方法值
	scale := p.ScaleBy
	fmt.Println(distanceFromP(q))
	scale(2)
	fmt.Println("-----------------------------")

	// 方法表达式
	distanceExp := (*Point).Distance
	scaleExp := (*Point).ScaleBy
	fmt.Println(distanceExp(&p, q))
	scaleExp(&q, 2)
}
