package main

import "Coder/code"

func main() {
	var coder code.ICoder = code.NewGopher("Go") // 用指针类型初始化接口，接口可调用接收者为指针类型的方法，也可调用接收者为值类型的方法；
	// var coder code.ICoder = code.CreateGopher("Go") // 用值类型初始化接口，接口只能调接收者为值类型的方法，那么结构体类型实现的方法其接收者只能是值类型；

	coder.Code()  // (coder)Code()
	coder.Debug() // (*coder)Debug()

	// 所以呢：结构体实现接口时（实现接口的方法），都使用指针类型作为其接收者；初始化接口时，使用指针类型初始化接口；
}
