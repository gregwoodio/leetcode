package problem0007

import "math"

func reverse(x int) int {
	output := 0
	isNegative := x < 0

	if isNegative {
		x *= -1
	}

	for x > 0 {
		output *= 10
		output += x % 10
		x /= 10
	}

	if isNegative {
		output *= -1
	}

	if output > math.MaxInt32 || output < math.MinInt32 {
		return 0
	}

	return output
}
