package Stringer

import (
	"testing"
)

func TestLower(t *testing.T) {
	var kata Words = "REZA"

	kata.Lower()

	if kata != "reza" {
		t.Error("words is wrong!")
	}
}
