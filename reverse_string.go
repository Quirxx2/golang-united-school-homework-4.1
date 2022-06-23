package reverse_string

func ReverseString(input string) (output string) {
	runes1 := []rune(input)
	var runes2 = ""
	for i := len(runes1) - 1; i > -1; i-- {
		runes2 = runes2 + string(runes1[i])
	}
	return runes2
}
