package main

import (
	"errors"
	"fmt"
	"strings"
)

type calculation func(int, int) int

func main() {
	fmt.Println("hello go")

	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))

	f := adder(10)
	fmt.Println(f(20))
	fmt.Println(f(30))

	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("dog"))
	fmt.Println(txtFunc("cat"))

	fmt.Println(f1()) // 5
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

	testDefer()
	testPanic()

	left := dispatchCoin()
	fmt.Println("剩下：", left)

	// 函数类型与变量
	var ad calculation = add
	fmt.Println(ad(1, 2))
	var su calculation = sub
	fmt.Println(su(5, 3))

	// 函数作为返回值
	addFunc, _ := do("+")
	fmt.Println(addFunc(1, 2))
}

func adder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// defer
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
} // 5

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
} // 6

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
} // 5

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
} // 5

// defer 的执行时机
// return语言不是原子操作：分为给返回值赋值和RET指令
// defer的执行，刚好在给返回值赋值操作后，在RET指令执行前
// 返回值赋值  defer的函数执行  RET指令执行
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// defer在注册其延时函数时，要确定其参数的值
// 后defer的函数先执行
func testDefer() {
	fmt.Println("testDefer")
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

func funcA() {
	fmt.Println("funcA")
}

func funcB() {
	defer func() { // recover() 在defer调用的函数中才有效
		err := recover()
		if err != nil {
			fmt.Println("recover from funcB")
		}
	}()
	panic("panic in funcB") // 在可能引发panic的语句之前定义defer
}

func funcC() {
	fmt.Println("funcC")
}

func testPanic() {
	funcA()
	funcB()
	funcC()
}

// 函数类型与变量
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func do(oper string) (func(int, int) int, error) {
	switch oper {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("oper is not recognition")
		return nil, err
	}
}
