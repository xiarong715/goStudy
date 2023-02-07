package person

import "fmt"

type Person struct {
	Name string
	Age  int8
}

func (p Person) GetName() string {
	return p.Name
}

func (p Person) GetAge() int8 {
	return p.Age
}

func (p Person) Print() {
	fmt.Printf("name = %s, age = %d\n", p.GetName(), p.GetAge())
}
