package Stringer

import (
	// "fmt"
	"strings"
)

func (w *Words) CamelCase(separator string) *Words {

	words := string(*w)

	splits := strings.Split(words, separator)

	var result string

	for a := 0; a < len(splits); a++ {

		looper := Words(splits[a])

		if a == 0 {
			looper.LowerFirst()
		}

		if a > 0 {
			looper.UpperFirst()
		}

		result = result + string(looper)
	}

	*w = Words(result)

	return w
}
