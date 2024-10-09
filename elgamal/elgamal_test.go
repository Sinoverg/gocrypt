package elgamal

import "testing"

func isIntSimpleFirstTest(t *testing.T) {
	have := int64(11)
	want := true
	msg := IsIntSimple(have)
	if msg != want {
		t.Fatalf("Error testing isIntSimple(11): want %t, but have %t", msg, want)
	}
}

func isIntSimpleSecondTest(t *testing.T) {
	have := int64(17)
	want := true
	msg := IsIntSimple(have)
	if msg != want {
		t.Fatalf("Error testing isIntSimple(17): want %t, but have %t", msg, want)
	}
}

func isIntSimpleThirdTest(t *testing.T) {
	have := int64(111)
	want := true
	msg := IsIntSimple(have)
	if msg != want {
		t.Fatalf("Error testing isIntSimple(111): want %t, but have %t", msg, want)
	}
}

func isIntSimpleFourTest(t *testing.T) {
	have := int64(111)
	want := true
	msg := IsIntSimple(have)
	if msg != want {
		t.Fatalf("Error testing isIntSimple(111): want %t, but have %t", msg, want)
	}
}
