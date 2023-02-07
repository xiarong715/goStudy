package main

import "fmt"

func main() {
	switch letter := 'a'; letter {
	case 'a':
		fmt.Printf("the letter is 'a'\n")
	case 'b':
		fmt.Printf("the letter is 'b'\n")
	case 'c':
		fmt.Printf("the letter is 'c'\n")
	default: // default 语句可以放到switch中的任意位置，并不影响switch的判断逻辑；
		fmt.Printf("the letter is %v", letter)
	}
	fmt.Printf("hello world")
}
