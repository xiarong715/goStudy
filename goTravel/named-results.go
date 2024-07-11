package main

import (
	"log"
	"testing"
)

func TestNamedReuslt(t *testing.T) {
	log.Println(split(17))
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
