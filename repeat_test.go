package Stringer

import (
	"testing"
)

func TestRepeat(t *testing.T) {

	var kata Words = "Reza"

	kata.Repeat(3)

	if kata != "RezaRezaReza" {
		t.Error("The words is wrong!")
	}
}
