package main

import "log"

// Shaper shaper
type Shaper interface {
	Area() float32
}

// Square square
type Square struct {
	side float32
}

// Rectangle ractangle
type Rectangle struct {
	length, width float32
}

func main() {
	square := &Square{5}
	rectangle := &Rectangle{2, 3}
	shapers := []Shaper{square, rectangle}

	for n := range shapers {
		log.Printf("图形数据:%v", shapers[n])
		log.Printf("图形面积:%v", shapers[n].Area())
	}

	var shaper1 Shaper
	var shaper2 Shaper
	shaper1 = square
	shaper2 = rectangle
	log.Println(shaper1.Area(), shaper2.Area())
}

// Area square area
func (s *Square) Area() float32 {
	return s.side * s.side
}

// Area rectangle area
func (r *Rectangle) Area() float32 {
	return r.length * r.width
}
