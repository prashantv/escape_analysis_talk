package main

import (
	"fmt"
	"io"
	"os"
)

var _output io.Writer = os.Stdout

type User struct {
	Name string
}

func main() {
	u := &User{Name: "Prashant"}
	display(u)
}

func display(u *User) {
	fmt.Fprintln(_output, u.Name)
}
