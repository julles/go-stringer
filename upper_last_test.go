package Stringer

import (
	"testing"
)

func TestUpperLast(t *testing.T) {
	var kata Words = "reza"

	kata.UpperLast()

	if kata != "rezA" {
		t.Error("words is wrong!")
	}
}
