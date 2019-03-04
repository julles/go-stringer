package Stringer

import (
	"testing"
)

func TestReplace(t *testing.T) {

	var words Words = "Reza"

	words.Replace("a", "o")

	if words != "Rezo" {
		t.Error("the words is wrong!")
	}

}
