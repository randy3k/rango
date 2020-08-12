package char

func ANSI(s string) Chars {
	// TODO: parse ANSI attributes
	runes := []rune(s)
	chars := make(Chars, len(runes))
	for i, r := range runes {
		chars[i] = NewChar(r, DefaultAttributes)
	}
	return chars
}
