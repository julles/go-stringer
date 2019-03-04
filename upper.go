package Stringer

import (
	"strings"
)

func (w *Words) Upper() *Words {
	*w = Words(strings.ToUpper(string(*w)))

	return w
}
