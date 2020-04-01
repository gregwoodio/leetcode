package problem0001

// naive implementation
// func twoSums(nums []int, target int) []int {
// 	for i, first := range nums {
// 		for j, second := range nums {
// 			if j == i {
// 				continue
// 			}

// 			if first+second == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}

// 	return []int{-1, -1}
// }

// implementation with map
func twoSums(nums []int, target int) []int {
	// map of value to index
	m := make(map[int]int)
	for i, val := range nums {
		m[val] = i
	}

	for i, val := range nums {
		j, ok := m[target-val]
		if i != j && ok {
			return []int{i, j}
		}
	}

	return []int{-1, -1}
}
