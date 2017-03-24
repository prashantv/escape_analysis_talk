package main

import (
	"io/ioutil"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	_output = ioutil.Discard

	for i := 0; i < b.N; i++ {
		main()
	}
}
