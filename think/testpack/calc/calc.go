package calc

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int64 {
	return (int64)(a * b)
}

func Div(a, b int) int {
	if b == 0 {
		panic("b == 0")
	}
	return a / b
}
