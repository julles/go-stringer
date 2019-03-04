package Stringer

import (
	"testing"
)

func TestWords(t *testing.T) {
	var words Words = "Reza"

	if words != "Reza" {
		t.Error("Words is wrong")
	}
}
