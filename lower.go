package Stringer

import (
	"strings"
)

func (w *Words) Lower() *Words {
	*w = Words(strings.ToLower(string(*w)))

	return w
}
