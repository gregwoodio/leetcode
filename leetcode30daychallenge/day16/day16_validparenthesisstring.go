package day16

func checkValidString(s string) bool {
	left, right := 0, 0

	for i, ch := range s {
		if ch == '(' || ch == '*' {
			left++
		} else {
			left--
		}

		rightCh := s[len(s)-i-1]
		if rightCh == ')' || rightCh == '*' {
			right++
		} else {
			right--
		}

		if left < 0 || right < 0 {
			return false
		}
	}

	return true
}
