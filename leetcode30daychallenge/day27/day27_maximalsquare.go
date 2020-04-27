package day27

func maximalSquare(matrix [][]byte) int {
	rows := len(matrix)
	if len(matrix) == 0 {
		return 0
	}

	cols := len(matrix[0])
	ans := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				ans = max(ans, helper(i, j, 1, rows, cols, matrix))
			}
		}
	}

	return ans
}

func helper(i, j, count, rows, cols int, matrix [][]byte) int {
	if !(check(i+1, rows) && check(j+1, cols)) {
		return (count) * (count)
	}

	for k := 0; k <= count; k++ {
		if matrix[i+1-k][j+1] == '0' || matrix[i+1][j+1-k] == '0' {
			return (count) * (count)
		}
	}

	return helper(i+1, j+1, count+1, rows, cols, matrix)
}

func check(n, length int) bool {
	return 0 <= n && n < length
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func maximalSquare(matrix [][]byte) int {
// 	max := 0

// 	for i := range matrix {
// 		for j := range matrix {
// 			if matrix[i][j] == '1' {
// 				current := expandSquare(i, j, matrix)

// 				if current > max {
// 					max = current
// 				}
// 			}
// 		}
// 	}

// 	return max
// }

// func expandSquare(i, j int, matrix [][]byte) int {
// 	size := 1

// 	for {
// 		for m := i; m < size+i && len(matrix[0]) > size+i; m++ {
// 			if matrix[m][size+i] != '1' {
// 				return (size - 1) * (size - 1)
// 			}
// 		}

// 		for n := j; n < size+j && len(matrix) > size+j; n++ {
// 			if matrix[size+j][n] != '1' {
// 				return (size - 1) * (size - 1)
// 			}
// 		}

// 		size++
// 	}
// }
