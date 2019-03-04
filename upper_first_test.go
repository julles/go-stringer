package Stringer

import (
	"testing"
)

func TestUpperFirst(t *testing.T) {
	var kata Words = "reza"

	kata.UpperFirst()

	if kata != "Reza" {
		t.Error("words is wrong!")
	}
}
