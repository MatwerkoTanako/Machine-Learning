package main

import (
	"errors"
	"fmt"
)

func Hello(name string) error {
	x := "hello world, "
	x += name
	fmt.Println(x)
	return errors.New("you suck at coding")
}

func main() {
	name := "Masaki"
	Hello(name)
}
