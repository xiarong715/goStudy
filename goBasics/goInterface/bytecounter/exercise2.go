package main

import (
	"fmt"
	"io"
)

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	return w, nil
}

func main() {
	fmt.Println("exercise2")
}
