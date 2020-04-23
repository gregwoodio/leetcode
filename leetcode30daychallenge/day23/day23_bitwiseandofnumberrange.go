package day23

func rangeBitwiseAnd(m int, n int) int {
	ands := m

	for i := m + 1; i <= n; i++ {
		ands &= i
	}

	return ands
}
