package main

import (
	"fmt"
	"os"
	"package/testconv"
	"strconv"
)

func main() {
	test1()
	fmt.Println("-----------------------")
	test2()
}

func test1() {
	c := testconv.Celsius(30)
	fmt.Println(c) // Implicitly call String
	fmt.Printf("\t%v\n", testconv.CToF(c))
	// fmt.Printf("\t%v\n", testconv.CToF(30))
	fmt.Printf("\t%v\n", testconv.CToK(c))
	// fmt.Println(c.String()) // Explicitly call String

	f := testconv.Fahrenheit(30)
	fmt.Println(f)
	// fmt.Println(c == f) // complier error, type mismatch
	fmt.Printf("\t%s\n", testconv.FToC(f))
	fmt.Printf("\t%s\n", testconv.FToK(f))

	k := testconv.Kelvin(30)
	fmt.Println(k)
	fmt.Printf("\t%v\n", testconv.KToC(k))
	fmt.Printf("\t%v\n", testconv.KToF(k))
}

func test2() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "parse float: %v", err)
			return
		}
		c := testconv.Celsius(t)
		f := testconv.Fahrenheit(t)
		k := testconv.Kelvin(t)
		fmt.Printf("%s = %s\t%s = %s\n%s = %s\t%s = %s\n%s= %s\t%s = %s\n",
			c, testconv.CToF(c), c, testconv.CToK(c),
			f, testconv.FToC(f), f, testconv.FToK(f),
			k, testconv.KToC(k), k, testconv.KToF(k))
	}
}
