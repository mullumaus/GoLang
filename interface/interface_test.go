package main

import (
	"fmt"
	"testing"
)

func describe(i interface{}) {
	fmt.Printf("%v %T\n", i, i)
}
func TestEmptyInterface(t *testing.T) {
	emptyInterface()
}

func TestNilInterface(t *testing.T) {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "Hello world"
	describe(i)
}
