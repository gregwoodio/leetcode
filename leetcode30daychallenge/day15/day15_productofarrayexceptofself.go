package day15

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	products := make([]int, len(nums))
	products[0] = 1

	for i := 1; i < len(nums); i++ {
		products[i] = nums[i-1] * products[i-1]
	}

	p := 1
	for i := len(nums) - 2; i >= 0; i-- {
		p *= nums[i+1]
		products[i] *= p
	}

	return products
}
