package day14

func stringShift(s string, shift [][]int) string {
	runes := []rune(s)
	l := len(runes)

	for _, currentShift := range shift {
		if len(currentShift) != 2 {
			panic("Malformed shift input")
		}

		dist := currentShift[1] % l

		if currentShift[0] == 0 { // left shift
			runes = append(runes[dist:], runes[:dist]...)
		} else { // right shift
			runes = append(runes[l-dist:], runes[:l-dist]...)
		}
	}

	return string(runes)
}
