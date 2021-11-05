package main

import "fmt"

func main() {
	fmt.Println("strStr")
	res := strStr("aabcde", "cde")
	fmt.Println("res = ", res)
}

func strStr(hey string, needle string) int {
	if len(needle) == 0 {
		return -1
	}
	// 遍历hey
	var i, j int
	for i = 0; i < len(hey)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			if hey[i+j] != needle[j] {
				break
			}
		}
		if j == len(needle) { //needle中的字符全部比较完
			return i
		}
	}
	return -1
}
