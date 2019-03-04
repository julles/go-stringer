package Stringer

import (
	"strings"
)

func (w *Words) Repeat(count int) *Words {

	*w = Words(strings.Repeat(string(*w), count))

	return w
}
