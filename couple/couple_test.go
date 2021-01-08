package couple_test

import (
	"prospersof/couple"
	"reflect"
	"testing"
)

func TestCouple1(t *testing.T) {
	given := "abcdef"
	want := []string{"ab", "cd", "ef"}
	get := couple.Couple(given)
	if !reflect.DeepEqual(want, get) {
		t.Errorf("given %v want %v but get %v\n", given, want, get)
	}
}

func TestCouple2(t *testing.T) {
	given := "abcdefg"
	want := []string{"ab", "cd", "ef", "g*"}
	get := couple.Couple(given)
	if !reflect.DeepEqual(want, get) {
		t.Errorf("given %v want %v but get %v\n", given, want, get)
	}
}
