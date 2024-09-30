package des

import (
	"slices"
	"testing"
)

func TestS1(t *testing.T) {
	have := []bool{false, false, false, false}
	want := []bool{true, false}
	msg := s1(have)
	if slices.Equal(msg, want) {
		t.Fatalf(`s1("0000") = %s, want match for %s`, PrintD(msg), PrintD(want))
	}
}
