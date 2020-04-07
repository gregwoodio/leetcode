package day07

func countElements(arr []int) int {
	elements := make(map[int]int)

	for _, val := range arr {
		elements[val]++
	}

	count := 0
	for _, val := range arr {
		if elements[val+1] > 0 {
			count++
		}
	}

	return count
}
