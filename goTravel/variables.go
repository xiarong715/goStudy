package main

import (
	"log"
	"testing"
)

var c, java, python bool

func TestVariables(t *testing.T) {
	var i int
	log.Println(i, c, java, python)
}
