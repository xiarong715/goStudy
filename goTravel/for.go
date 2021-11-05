package main

import "log"

func main() {
	for i := 10; i > 0; i-- {
		log.Print(i)
	}

	sum := 1
	for sum < 100 {
		sum += sum
	}
	log.Print(sum)

	// infinite loop
	for {

	}
}
