package Stringer

import (
	"testing"
)

func TestSubstring(t *testing.T) {
	var kata Words = "muhamad reza"

	kata.Substring(2, 5)

	if kata != "ham" {
		t.Error("words is wrong!")
	}
}
