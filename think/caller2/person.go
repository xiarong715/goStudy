package main

type Person struct {
	Name string
	Age  int8
}

// 值接收者
func (p Person) getName() string {
	return p.Name
}

// 值接收者
func (p Person) getAge() int8 {
	return p.Age
}

// 指针接收者
func (p *Person) setName(name string) {
	p.Name = name
}

// 指针接收者
func (p *Person) setAge(age int8) {
	p.Age = age
}
