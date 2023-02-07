package main

import "fmt"

func main() {
	per := Person{
		Name: "tom",
		Age:  18,
	}
	stu := Student{
		Person: Person{
			Name: "jim",
			Age:  18,
		},
		ID: 1,
	}

	fmt.Printf("name = %s, age = %d\n", per.getName(), per.getAge())
	fmt.Printf("name = %s, age = %d, id = %d\n", stu.getName(), stu.getAge(), stu.getID())

	//值接收者，还是指针接收者的方法（非接口方法），可用值接收者调用，也可用指针接收者调用。因为编译器会自动转化。
	per.setName("sot")
	stu.setName("dim")

	fmt.Printf("name = %s, age = %d\n", per.getName(), per.getAge())
	fmt.Printf("name = %s, age = %d, id = %d\n", stu.getName(), stu.getAge(), stu.getID())
}
