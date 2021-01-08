package main

import (
	"fmt"
	"strconv"
)

// Int represent integer
type Int int

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

// Set set new value
func (i *Int) Set(n int) {
	*i = Int(n)
}

// Int return int value
func (i Int) Int() int {
	return int(i)
}

func main() {
	var i Int = 99
	fmt.Printf("%d string is %q\n", i, i.String())

	fmt.Println("i.Int() is", i.Int())

	i.Set(100)
	fmt.Println("i =", i)
}
