package main

import "io"

var _output io.Writer

func main() {
	bs := []byte("Prashant")
	emptyWrite(bs)
}

func emptyWrite(bs []byte) {}
