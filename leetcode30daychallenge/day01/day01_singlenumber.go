package day01

// Given a non-empty array of integers, every element appears twice except for one. Find that single one.

// Note:

// Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

// Example 1:

// Input: [2,2,1]
// Output: 1

// Example 2:

// Input: [4,1,2,1,2]
// Output: 4

func singleNumber(nums []int) int {
	myMap := make(map[int]bool)

	for _, num := range nums {
		_, ok := myMap[num]
		if !ok {
			myMap[num] = false
		} else {
			myMap[num] = true
		}
	}

	for key, val := range myMap {
		if val == false {
			return key
		}
	}

	return -1
}

// Better solution (not mine):
// func singleNumber(nums []int) int {
//     result:=0
//     for _,v:= range nums{
//        result^=v
//     }
//     return result
// }
