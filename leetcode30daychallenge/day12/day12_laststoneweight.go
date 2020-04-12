package day12

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		maxIndex := indexOfMax(stones)
		y := stones[maxIndex]
		stones = append(stones[0:maxIndex], stones[maxIndex+1:]...)

		maxIndex = indexOfMax(stones)
		x := stones[maxIndex]
		stones = append(stones[0:maxIndex], stones[maxIndex+1:]...)

		if x != y {
			stones = append(stones, y-x)
		}
	}

	if len(stones) > 0 {
		return stones[0]
	}

	return 0
}

func indexOfMax(nums []int) int {
	maxVal, maxIndex := nums[0], 0

	for i, val := range nums {
		if val > maxVal {
			maxIndex = i
			maxVal = val
		}
	}

	return maxIndex
}
