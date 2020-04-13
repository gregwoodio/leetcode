package day13

func findMaxLength(nums []int) int {
	max := 0

	// map of count to last seen index with that count
	countMap := make(map[int]int)
	countMap[0] = -1

	count := 0

	for i, num := range nums {
		if num == 0 {
			count--
		} else if num == 1 {
			count++
		}

		if j, ok := countMap[count]; ok {
			current := i - j
			if current > max {
				max = current
			}
		} else {
			countMap[count] = i
		}
	}

	return max
}
