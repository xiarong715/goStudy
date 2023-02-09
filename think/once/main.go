package main

import (
	"fmt"
	singleton "once/singlet" // 文件夹就是包，当包名与文件夹名不同时，导入包时需要带上包名
)

func main() {
	inst1 := singleton.GetInstance()
	inst2 := singleton.GetInstance()
	if inst1 == inst2 {
		fmt.Println("inst1 == inst2")
	}
	fmt.Println("hello once")
}
