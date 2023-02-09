package snow

import (
	"fmt"
	"testpack/calc"
)

func TestCalc() {
	fmt.Println(calc.Add(1, 2))
	fmt.Println(calc.Sub(1, 2))
	fmt.Println(calc.Mul(1, 2))
	fmt.Println(calc.Div(1, 2))
}
