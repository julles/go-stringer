package Stringer

import (
	"strings"
)

func (w *Words) UpperLast() *Words {
	splits := strings.Split(string(*w), "")

	appends := ""
	for a := 0; a < len(splits); a++ {

		result := splits[a]

		if a+1 == len(splits) {
			result = strings.ToUpper(result)
		}

		appends += result
	}

	*w = Words(appends)

	return w
}
