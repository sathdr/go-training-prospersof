package fizzbuzz_test

import (
	"math/rand"
	"prospersof/fizzbuzz"
	"testing"
)

func TestFizzBuzzGiven1(t *testing.T) {
	given := 1
	want := "1"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFizzBuzzGiven2(t *testing.T) {
	given := 2
	want := "2"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFizzBuzzGiven3(t *testing.T) {
	given := 3
	want := "Fizz"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFizzBuzzGiven4(t *testing.T) {
	given := 4
	want := "4"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFizzBuzzGiven5(t *testing.T) {
	given := 5
	want := "Buzz"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFizzBuzzGiven6(t *testing.T) {
	given := 6
	want := "Fizz"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFizzBuzzGiven7(t *testing.T) {
	given := 7
	want := "7"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFizzBuzzGiven8(t *testing.T) {
	given := 8
	want := "8"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFizzBuzzGiven9(t *testing.T) {
	given := 9
	want := "Fizz"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFizzBuzzGiven10(t *testing.T) {
	given := 10
	want := "Buzz"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFizzBuzzGiven12(t *testing.T) {
	given := 12
	want := "Fizz"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFizzBuzzGiven15(t *testing.T) {
	given := 15
	want := "FizzBuzz"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFizzBuzzGiven18(t *testing.T) {
	given := 18
	want := "Fizz"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFizzBuzzGiven20(t *testing.T) {
	given := 20
	want := "Buzz"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFizzBuzzGiven30(t *testing.T) {
	given := 30
	want := "FizzBuzz"

	get := fizzbuzz.Say(given)
	if want != get {
		t.Errorf("given %d want %q but get %q\n", given, want, get)
	}
}

func TestFormat(t *testing.T) {
	want := "Fizz-Fizz-Fizz-Fizz"

	given := func() int {
		return 3
	}

	get := fizzbuzz.Format(given)

	if want != get {
		t.Errorf("given %v want %q but get %q\n", given(), want, get)
	}
}

func xTestFormatInReal(t *testing.T) {
	r := rand.New(rand.NewSource((1)))

	want := "Fizz-Fizz-Fizz-Fizz"

	given := func() int {
		return r.Intn(100) + 1
	}

	get := fizzbuzz.Format(given)

	if want != get {
		t.Errorf("given %v want %q but get %q\n", given(), want, get)
	}
}
func TestFormatSlices(t *testing.T) {
	numbers := []int{1, 2, 3, 4}

	want := "1-2-Fizz-4"

	given := func() int {
		defer func() { numbers = numbers[1:] }()
		return numbers[0]
	}

	get := fizzbuzz.Format(given)

	if want != get {
		t.Errorf("given %v want %q but get %q\n", given(), want, get)
	}
}

type Intn struct {
	i       int
	numbers []int
}

func (x *Intn) Intn(n int) int {
	defer func() { x.i++ }()
	return x.numbers[x.i] - 1
}

func TestFormatInterface(t *testing.T) {
	given := &Intn{numbers: []int{1, 3, 5, 15}}

	want := "1-Fizz-Buzz-FizzBuzz"

	get := fizzbuzz.FormatGetInterface(given)

	if want != get {
		t.Errorf("given %v want %q but get %q\n", given, want, get)
	}
}

func TestFormatByClosure(t *testing.T) {
	given := []int{1, 3, 5, 15}

	want := "1-Fizz-Buzz-FizzBuzz"

	closure := func(numbers []int) func() int {
		i := 0
		return func() int {
			defer func() { i++ }()
			return numbers[i]
		}
	}

	get := fizzbuzz.Format(closure(given))

	if want != get {
		t.Errorf("given %v want %q but get %q\n", given, want, get)
	}
}

func (fn fizzbuzz.RandFunc) Intn(n int) int {
	return fn(n)
}
