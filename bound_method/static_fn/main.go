package main

import (
	"io"
	"os"
	"time"
)

var _output io.Writer = os.Stdout

func printHelloWorld() {
	io.WriteString(_output, "Hello World\n")
}

func main() {
	time.AfterFunc(time.Millisecond, printHelloWorld)
}
