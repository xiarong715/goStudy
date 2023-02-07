package animal

import "fmt"

type Animal struct {
	Name string
	Age  int8
}

type Behavior interface {
	Yell()
	Sleep()
}

func (animal *Animal) Yell() {
	fmt.Printf("%s (animal) yell\n", animal.Name)
}

func (animal *Animal) Sleep() {
	fmt.Printf("%s (animal) sleep\n", animal.Name)
}

func (animal *Animal) SetName(name string) {
	animal.Name = name
}

func (animal *Animal) SetAge(age int8) {
	animal.Age = age
}

func Yell(behavior Behavior) {
	behavior.Yell()
}

func Sleep(behavior Behavior) {
	behavior.Sleep()
}
