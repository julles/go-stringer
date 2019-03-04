package Stringer

import (
	"strings"
)

func (w *Words) Replace(search string, replace string) *Words {

	*w = Words(strings.Replace(string(*w), search, replace, 1))

	return w
}
