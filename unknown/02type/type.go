package main

import (
	"fmt"
	"math"
)

// go 类型：
// 值类型：数值类型、字符串类型、布尔类型、数组类型、结构体类型；
// 引用类型：函数类型、指针类型、切片类型、Map类型、Chan类型、接口类型。

// 数值类型：整型、浮点型、复数型。
// 整型：int/uint int8 (byte) /uint8 int16/uint16 int32 (rune) /uint32
// 浮点型：float32 float64
// 复数型：complex64（实部虚部各占32） complex128（实部虚部各占64）

// 字符串类型：string
// 布尔类型：bool
// 数组类型：[n]type  如：[2]int，有两个int元素的数组
// 结构体类型：struct {}

// 函数类型： func funcname(params) {}
// 指针类型： *T，T为类型，如：*int，指向值为int类型的地址空间
// 切片类型： []T，T为类型
// Map类型： map[T1][T2]，T1,T2为类型，可相同也可不同
// Chan类型： chan T，T为类型，如：chan int，int型的chan
// 接口类型：interface {}

// 类型
// bool   1字节  true false
// int/unit   32/64位，根据平台位数决定
// int8 (byte) /unit8 1字节 -128~127   0~255
// int16/uint16  2字节
// int32 (rune) /uint32  4字节	-2^32/2~2^32/2 -1    0~2^32-1

// 复数类型 complex64/complex128   8/16字节

// 指针类型 uintptr

// 其他类型 arry struct string

// 引用类型 slice map chan（用于并发时的通信）

// 接口类型 interface
// 函数类型 func

// byte (uint8的别名)
// rune (int32的别名)

// 类型的零值
// 一般类型的零值为 0
// bool 的零值为 false
// string 的零值为 ""
// 引用类型的零值为 nil

// 类型的最大值与最小值
// math.Maxint16
// math.Minint16

// 类型别名
type (
	byte  = int8
	rune  = int32
	myInt = int
)

// 自定义类型
type (
	person struct{}
	newInt int
)

// 变量的声明与赋值
// 变量的声明  var <变量名称> <变量类型>
// 变量的赋值格式 <变量名称> = <表达式>
// 变量声明并赋值 var <变量名称> [变量类型] = <表达式>， 此时可省变量名

// _ 忽略符

// 变量类型转换，go中只能显示转换，不同类型之间不能转换，如 int ---> bool

func main() {
	var b bool
	fmt.Println(b)

	var str string
	fmt.Printf("str = %s\n", str)

	var i interface{}
	fmt.Println(i)

	var n int
	fmt.Println(n)

	var p uintptr // 指针
	fmt.Println(p)

	var arry [2]int   // 数组
	fmt.Println(arry) // [0 0]

	var s []int    // 切片
	fmt.Println(s) // []

	fmt.Println(math.MinInt8)

	var a int // 先声明，后赋值
	a = 1

	d := 1 // 声明并赋值，只可用在函数中
	fmt.Printf("%d %d\n", a, d)

	var c = 3 // 自动推断类型
	fmt.Printf("%T\n", c)

	f := float32(c)
	// bo := bool(c)
	fmt.Println(f)

	var in = 65
	st := string(in) // 转化为字符
	fmt.Println(st)
}
