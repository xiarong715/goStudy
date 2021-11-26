package main

import "fmt"

var nums []int = []int{1, 1, 2, 2, 3}

func main() {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	fmt.Println(res)
}
