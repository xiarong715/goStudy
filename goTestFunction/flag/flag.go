package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

// ./flag 5s
// ./flag 200ms

func main() {
	fmt.Println("flag")
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
