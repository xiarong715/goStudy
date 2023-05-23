package main

import "fmt"

func main() {
	msg := "abc"
	json := `{"reply":"ok", "result":"` + msg + `"}` // 用 `` 号组合原生的字符串，非常好用
	msg1 := `{"reply":"failed", "result":"` + msg + `"}`
	fmt.Println(json)
	fmt.Println(msg1)
}
