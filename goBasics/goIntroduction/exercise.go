package main

import (
	"fmt"
	"os"
)

// print os.Args

func test1() {
	// var s, sep string
	// for i := 0; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)

	var s, sep string
	for _, arg := range os.Args[:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	// fmt.Println(strings.Join(os.Args[:], " "))

	// fmt.Println(os.Args[:])
}

func test2() {
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}
}

func main() {
	fmt.Println("go prgram struct exercise")
	fmt.Println("------------------------------")
	test1()
	fmt.Println("------------------------------")
	test2()
}
