package fizzbuzz_test

import (
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
