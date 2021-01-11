package fizzbuzz

import (
	"fmt"
	"strconv"
)

// Say return a number
func Say(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}
	if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}

// RandFunc is a function
type RandFunc func(int) int

// Format format random 4 number fizzbuzz
func Format(randFunction func() int) string {
	return fmt.Sprintf("%s-%s-%s-%s", Say(randFunction()), Say(randFunction()), Say(randFunction()), Say(randFunction()))
}

func FormatGetInterface(random interface{ Intn(n int) int }) string {
	return fmt.Sprintf("%s-%s-%s-%s", Say(random.Intn(100)+1), Say(random.Intn(100)+1), Say(random.Intn(100)+1), Say(random.Intn(100)+1))
}
