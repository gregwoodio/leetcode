package day25

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	canJump := make([]bool, len(nums))
	canJump[len(canJump)-1] = true

	for j := len(nums) - 2; j >= 0; j-- {
		num := nums[j]

		for i := j + 1; i < j+1+num && i < len(nums); i++ {
			if canJump[i] == true {
				canJump[j] = true
			}
		}
	}

	return canJump[0]
}
