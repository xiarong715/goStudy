package main

import (
	"log"
	"testing"
)

// 打印变量的类型

func TestTypeInference(t *testing.T) {
	v := true
	log.Printf("v is of type %T", v)
}
