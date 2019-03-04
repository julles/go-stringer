package Stringer

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var words Words = "reza"

	words.Reverse()

	if words != "azer" {
		t.Error("Words is wrong!")
	}
}
