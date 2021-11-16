package geometry

import "math"

type Point struct{ X, Y float64 }

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func (p Path) Distance() float64 {
	// 法1
	// path := 0.0
	// for i := 0; i < len(p)-1; i++ {
	// 	path += p[i].Distance(p[i+1])
	// }
	// return path

	// 法2
	path := 0.0
	for i := range p {
		if i > 0 {
			path += p[i-1].Distance(p[i])
		}
	}
	return path
}

func PathDistance(p Path) float64 {
	sum := 0.0
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}
