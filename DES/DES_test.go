package des

import (
	"slices"
	"testing"
)

func TestS1WithFFFF(t *testing.T) {
	have := []bool{false, false, false, false}
	want := []bool{false, true}
	msg := s1(have)
	if !slices.Equal(msg, want) {
		t.Fatalf(`s1("0000") = %s, want match for %s`, PrintD(msg), PrintD(want))
	}
}

func TestS1WithFTTF(t *testing.T) {
	have := []bool{false, true, true, false}
	want := []bool{true, false}
	msg := s1(have)
	if !slices.Equal(msg, want) {
		t.Fatal(`s1("0110") =" `, msg, `" %s, want "10" match for `, want)
	}
}

func TestS1WithTTTF(t *testing.T) {
	have := []bool{true, true, true, false}
	want := []bool{true, true}
	msg := s1(have)
	if !slices.Equal(msg, want) {
		t.Fatalf(`s1("1110") = %s, want match for %s`, PrintD(msg), PrintD(want))
	}
}

func TestS2WithFFFF(t *testing.T) {
	have := []bool{false, false, false, false}
	want := []bool{false, true}
	msg := s2(have)
	if !slices.Equal(msg, want) {
		t.Fatalf(`s2("0000") = %s, want match for %s`, PrintD(msg), PrintD(want))
	}
}

func TestS2WithFTTF(t *testing.T) {
	have := []bool{false, true, true, false}
	want := []bool{true, true}
	msg := s2(have)
	if !slices.Equal(msg, want) {
		t.Fatal(`s2("0110") =" `, msg, `" %s, want "10" match for `, want)
	}
}

func TestS2WithTTTF(t *testing.T) {
	have := []bool{true, true, true, false}
	want := []bool{false, false}
	msg := s2(have)
	if !slices.Equal(msg, want) {
		t.Fatalf(`s2("1110") = %s, want match for %s`, PrintD(msg), PrintD(want))
	}
}
func TestBToInt3(t *testing.T) {
	have := []bool{true, true}
	want := 3
	msg := bToInt(have)
	if msg != want {
		t.Fatalf(`bToInt("11") = %d, want match for %d`, msg, want)
	}
}
func TestBToInt1(t *testing.T) {
	have := []bool{false, true}
	want := 1
	msg := bToInt(have)
	if msg != want {
		t.Fatalf(`bToInt("01") = %d, want match for %d`, msg, want)
	}
}
func TestBToInt0(t *testing.T) {
	have := []bool{false, false}
	want := 0
	msg := bToInt(have)
	if msg != want {
		t.Fatalf(`bToInt("00") = %d, want match for %d`, msg, want)
	}
}
func TestBToInt4(t *testing.T) {
	have := []bool{true, false, false}
	want := 4
	msg := bToInt(have)
	if msg != want {
		t.Fatalf(`bToInt("100") = %d, want match for %d`, msg, want)
	}
}
