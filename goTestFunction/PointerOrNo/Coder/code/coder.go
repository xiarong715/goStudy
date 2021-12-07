package code

import "fmt"

type ICoder interface {
	Code()
	Debug()
}

type Gopher struct {
	lang string
}

func NewGopher(lang string) *Gopher {
	return &Gopher{lang: lang}
}

func CreateGopher(lang string) Gopher {
	return Gopher{lang: lang}
}

func (g *Gopher) Code() {
	fmt.Printf("I'm coding %s language\n", g.lang)
}

// 测试：接口用指针类型和值类型初始化的区别
// 结论：都使用指针类型作为接收者
func (g Gopher) Debug() {
	fmt.Printf("I'm debuging %s language\n", g.lang)
}
