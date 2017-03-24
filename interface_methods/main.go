package main

import (
	"io"
	"io/ioutil"
)

var _output io.Writer

func main() {
	// ioutil.Discard is a io.Writer
	// that throws away bytes.
	bs := []byte("Prashant")
	ioutil.Discard.Write(bs)
}
