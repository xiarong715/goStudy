package main

import (
	"fmt"
	"log"
	"math"
	"testing"
)

func TestIf(t *testing.T) {
	log.Println(sqrt(2), sqrt(-4))
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
