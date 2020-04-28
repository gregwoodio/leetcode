package day28

type FirstUnique struct {
	nums               []int
	counts             map[int]int
	highestUniqueIndex int
}

func Constructor(nums []int) FirstUnique {
	uniq := FirstUnique{
		nums:               make([]int, 0),
		counts:             make(map[int]int),
		highestUniqueIndex: 0,
	}

	for _, val := range nums {
		uniq.Add(val)
	}

	return uniq
}

func (this *FirstUnique) ShowFirstUnique() int {
	for ; this.highestUniqueIndex < len(this.nums); this.highestUniqueIndex++ {
		if this.counts[this.nums[this.highestUniqueIndex]] == 1 {
			return this.nums[this.highestUniqueIndex]
		}
	}

	return -1
}

func (this *FirstUnique) Add(value int) {
	this.nums = append(this.nums, value)
	this.counts[value]++
}

/**
 * Your FirstUnique object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.ShowFirstUnique();
 * obj.Add(value);
 */
