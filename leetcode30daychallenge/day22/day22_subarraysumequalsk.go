package day22

func subarraySum(nums []int, k int) int {
	sums := make([]int, len(nums))

	for i := 0; i < len(sums); i++ {
		sums[i] = nums[i]
		if i > 0 {
			sums[i] += sums[i-1]
		}
	}

	count := 0

	for i := 0; i < len(sums); i++ {
		if sums[i] == k {
			count++
		}

		for j := i + 1; j < len(sums); j++ {
			curr := sums[j] - sums[i]

			if curr == k {
				count++
			}
		}
	}

	return count
}
