package main

import "fmt"

func main() {
	fmt.Println("hello go")
	testSlice()
}

func testSlice() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	s := a[1:3]
	fmt.Printf("s:%v, len(s):%v, cap(s):%v\n", s, len(s), cap(s)) // [2,3], len(s):2, cap(s):4
	s2 := s[2:4]
	fmt.Printf("s2:%v, len(s2):%v, cap(s2):%v\n", s2, len(s2), cap(s2)) // [4 5], len(s):2, cap(s):2
}
