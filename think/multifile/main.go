package main

func Print(printer Printer) {
	printer.Print()
}

func main() {
	per := Person{
		name: "smith",
		age:  18,
	}
	stu := Student{
		Person: Person{
			name: "tom",
			age:  20,
		},
		id: 1,
	}
	Print(per)
	Print(stu)
}
