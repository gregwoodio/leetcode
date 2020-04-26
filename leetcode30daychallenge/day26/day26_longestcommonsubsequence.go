package day26

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return lcs(text1, text2, 0, 0, &dp)
}

func lcs(text1, text2 string, m, n int, dp *[][]int) int {
	if m == len(text1) || n == len(text2) {
		return 0
	}

	if (*dp)[m][n] != -1 {
		return (*dp)[m][n]
	}

	if text1[m] == text2[n] {
		(*dp)[m][n] = 1 + lcs(text1, text2, m+1, n+1, dp)
	} else {
		(*dp)[m][n] = max(lcs(text1, text2, m+1, n, dp), lcs(text1, text2, m, n+1, dp))
	}

	return (*dp)[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
