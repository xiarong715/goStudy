package main

import "fmt"

type Person struct {
	name   string
	age    int8
	dreams []string
}

// 因为map、slice是引用类型，在赋值时，使用深拷贝的方式
func (p *Person) SetDreams(dreams []string) {
	// p.dreams = dreams  // 浅拷贝

	// 深拷贝
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

func main() {
	p := Person{name: "小王子", age: 18} //不完全赋值，则指明字段赋值
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p.SetDreams(data)
	data[1] = "不睡觉"
	fmt.Println(p.dreams)
}
