package main

import (
	"io"
	"os"
	"time"
)

var _output io.Writer = os.Stdout

type S struct {
	s string
}

func (s S) Print() {
	io.WriteString(_output, s.s)
}

func main() {
	s := S{"Hello World\n"}
	time.AfterFunc(time.Millisecond, s.Print)
}
