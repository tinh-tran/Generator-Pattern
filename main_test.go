package main

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	for i:= range Count(1,99) {
		fmt.Println("Pass it around, put one up,", i, "bottles of beer on the wall")
	}
}
