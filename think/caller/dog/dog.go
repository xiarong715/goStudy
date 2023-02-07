package dog

import (
	"caller/animal"
	"fmt"
)

type Dog struct {
	animal.Animal
}

func (d *Dog) SetName(name string) {
	d.Animal.SetName(name)
}

func (d *Dog) SetAge(age int8) {
	d.Animal.SetAge(age)
}

func (d *Dog) Yell() {
	fmt.Printf("%s (cat) miaomiao\n", d.Animal.Name)
}

func (d *Dog) Sleep() {
	fmt.Printf("%s (cat) sleep\n", d.Animal.Name)
}
