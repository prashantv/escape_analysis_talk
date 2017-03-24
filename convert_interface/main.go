package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var _output io.Writer = os.Stdout

func main() {
	t := time.Now()

	// time.Time has a String() method, so can be passed as a Stringer.
	printStringer(t)
}

func printStringer(s fmt.Stringer) {
	io.WriteString(_output, s.String())
}
