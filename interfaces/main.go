package main

import "fmt"

type x struct{}

func (x) String() string {
	return "I'm String"
}

type Int int

func (Int) String() string {
	return "I'm Integer"
}

func PrintString(s fmt.Stringer) {
	fmt.Println(s.String())
}

func main() {
	PrintString(x{})
	PrintString(Int(1))
}
