package Stringer

func (w *Words) Substring(first int, last int) *Words {

	*w = "Ok"

	words := "Muhamad Reza Abdul Rohim"

	slice := []rune(words)

	theSubstring := string(slice[first:last])

	*w = Words(theSubstring)

	return w

}
