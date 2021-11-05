package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("./plugin/testPlugin.so")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("Hello")
	if err != nil {
		panic(err)
	}
	f.(func())()
	fmt.Println("main")
}
