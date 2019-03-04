package Stringer

import (
	"testing"
)

func TestLowerLast(t *testing.T) {
	var kata Words = "REZA"

	kata.LowerLast()

	if kata != "REZa" {
		t.Error("words is wrong!")
	}
}
