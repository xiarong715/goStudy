package main

import "log"

func main() {
	// bool
	// string
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte // uint8 的别名
	// rune // int32 的别名
	// float32 float64
	// complex64 complex128
	// int uint uintptr 在32位宽的系统上通常为32位宽，在64位宽的系统上通常为64位宽；
	// 当你需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型；

	var MaxInt uint64 = 1<<64 - 1
	log.Println(MaxInt)
}
