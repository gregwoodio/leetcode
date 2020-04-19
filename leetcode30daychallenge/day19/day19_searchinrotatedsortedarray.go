package day19

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	pivot := findPivot(nums)

	sorted := append(nums[pivot:], nums[:pivot]...)

	index := binarySearch(sorted, target, 0, len(sorted)-1)
	if index == -1 {
		return index
	}

	return (index + pivot) % len(nums)
}

func findPivot(nums []int) int {
	pivot := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			pivot++
		} else {
			break
		}
	}

	// the above logic gets us up to the highest valued index,
	// we need one more step to get the original start
	pivot++

	return pivot
}

func binarySearch(nums []int, target, low, high int) int {
	if high < low {
		return -1
	}

	mid := (low + high) / 2
	if nums[mid] > target {
		return binarySearch(nums, target, low, mid-1)
	} else if nums[mid] < target {
		return binarySearch(nums, target, mid+1, high)
	} else {
		return mid
	}
}
