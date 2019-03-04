package Stringer

import (
	"testing"
)

func TestUpper(t *testing.T) {
	var kata Words = "reza"

	kata.Upper()

	if kata != "REZA" {
		t.Error("words is wrong!")
	}
}
