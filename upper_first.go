package Stringer

import (
	"strings"
)

func (w *Words) UpperFirst() *Words {
	splits := strings.Split(string(*w), "")

	appends := ""

	for a := 0; a < len(splits); a++ {

		result := splits[a]

		if a == 0 {
			result = strings.ToUpper(result)
		}

		appends += result
	}

	*w = Words(appends)

	return w
}
