package main

import (
	"fmt"
	"runtime"
)

func TestSwitch() {
	fmt.Print("Go run on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}
