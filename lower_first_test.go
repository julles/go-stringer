package Stringer

import (
	"testing"
)

func TestLowerFirst(t *testing.T) {
	var kata Words = "REZA"

	kata.LowerFirst()

	if kata != "rEZA" {
		t.Error("words is wrong!")
	}
}
