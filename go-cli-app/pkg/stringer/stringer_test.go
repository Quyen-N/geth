package stringer

import (
	"testing"
)

func TestReverse(t *testing.T) {
	input := "Hello, 李祖阳"
	want := "阳祖李 ,olleH"
	have := Reverse(input)
	if have != want {
		t.Error("want:", want, ", have:", have)
	}
}

func TestInspect(t *testing.T) {
	// inspect digit
	input := "012a3c45b6"
	want_count := 7
	want_kind := "digit"
	have_count, have_kind := Inspect(input, true)
	if want_count != have_count || want_kind != have_kind {
		t.Error("inspect digit fail")
	}

	// inspect char
	input = "Hello"
	want_count = 5
	want_kind = "char"
	have_count, have_kind = Inspect(input, false)
	if want_count != have_count || want_kind != have_kind {
		t.Error("inspect char fail")
	}
}
