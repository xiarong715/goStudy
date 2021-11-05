package main

import (
	"log"
	"math"
)

func main() {
	log.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

// if语句中定义的变量，在if和else中都可使用；

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		log.Printf("%g >= %g", v, lim)
	}
	return lim
}
