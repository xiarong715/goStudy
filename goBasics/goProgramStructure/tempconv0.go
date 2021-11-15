package main

import "fmt"

type Fahrenheit float64
type Celsius float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	c := Celsius(30)
	fmt.Println(CToF(c))

	// f := CToF(c)
	// fmt.Println(c == f) // complier error, type mismatch
	fmt.Println(c.String())
	fmt.Println(c)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32) // 类型传换，不是函数调用，把c*9/5 + 32计算出的值传换为Fahrenheit类型，只会改变值类型，不会改变值本身
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g ℃", c)
}
