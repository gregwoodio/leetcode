package day09

func backspaceCompare(S string, T string) bool {
	return typeString(S) == typeString(T)
}

func typeString(str string) string {
	chars := []rune{}

	for _, ch := range str {
		if ch == '#' {
			// backspace character
			if len(chars) > 0 {
				chars = chars[0 : len(chars)-1]
			} else {
				chars = []rune{}
			}
		} else {
			// regular character
			chars = append(chars, ch)
		}
	}

	return string(chars)
}
