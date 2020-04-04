package day04

func moveZeroes(nums []int) {
	l, zeroes, idx := len(nums), 0, 0
	numsCopy := make([]int, l)
	copy(numsCopy, nums)

	for _, val := range numsCopy {
		if val == 0 {
			nums[l-1-zeroes] = 0
			zeroes++
		} else {
			nums[idx] = val
			idx++
		}
	}
}
