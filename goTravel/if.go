package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	log.Println(sqrt(2), sqrt(-4))
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
