package main

import (
	"fmt"
	"testing"
	"time"
)

func TestSwitchWithNoCondition(t *testing.T) {
	tm := time.Now()
	switch {
	case tm.Hour() < 12:
		fmt.Println("Good morning!")
	case tm.Hour() < 17:
		fmt.Println("Good aternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
