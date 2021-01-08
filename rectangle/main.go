package main

import "fmt"

// Rectangle represents a rectangle
type Rectangle struct {
	Width  float64
	Length float64
}

// Area find area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

func main() {
	r := Rectangle{Width: 10, Length: 6}
	fmt.Println(r.Area())
}
