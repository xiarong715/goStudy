package main

import (
	"fmt"
	"os"
)

type TerminalOut struct {
	DeviceName *os.File
}

func (t *TerminalOut) Write(msg string) {
	fmt.Fprintf(t.DeviceName, "%s", msg)
}
