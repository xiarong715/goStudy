package main

import (
	"log"
	"math"
	"testing"
)

func TestIfElse(t *testing.T) {
	log.Println(
		power(3, 2, 10),
		power(3, 3, 20),
	)
}

// if语句中定义的变量，在if和else中都可使用；

func power(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		log.Printf("%g >= %g", v, lim)
	}
	return lim
}
