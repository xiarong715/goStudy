package main

func main() {
	per := Person{
		Name: "hom",
		Age:  20,
	}

	stu := Student{
		Person: Person{
			Name: "tom",
			Age:  22,
		},
		ID: 1,
	}

	per.Print() // 指针接收者实现Print()，值调用者调用Print()没问题

	// var x Printer = per    // 指针类型接收者实现接口，则该值类型不能被赋给实现的接口。值类型接收者实现接口，则值类型或指针类型变量都可赋值给实现的接口。
	var x Printer = &per // 指针类型接收者实现接口，则该指针类型能被赋给实现的接口。
	x.Print()            // 指针类型调用指针类型实现的接口的方法

	x = &stu
	x.Print()

	// 需要用接口作为参数，实现多态。指针类型接收者实现接口，只有结构体的指针类型变量可被赋给接口变量。
	// 所有实现接口的类型，都能当作同一类型，都可作为同一类型处理。

	Print(&per)
	Print(&stu)
}
