package cat

import (
	"caller/animal"
	"fmt"
)

type Cat struct {
	animal.Animal
}

func (a *Cat) SetName(name string) {
	a.Animal.SetName(name)
}

func (a *Cat) SetAge(age int8) {
	a.Animal.SetAge(age)
}

func (a *Cat) Yell() {
	fmt.Printf("%s (cat) miaomiao\n", a.Animal.Name)
}

func (a *Cat) Sleep() {
	fmt.Printf("%s (cat) sleep\n", a.Animal.Name)
}
