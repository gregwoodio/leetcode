package day29

import (
	"math"
	"testing"
)

func TestMaxPath(t *testing.T) {
	nums := []int{1, 2, math.MinInt32, math.MinInt32, 3, math.MinInt32, math.MinInt32}
	root := makeTree(&nums)

	actual := maxPathSum(root)
	expected := 6

	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func makeTree(nums *[]int) *TreeNode {
	if len(*nums) <= 0 || (*nums)[0] == math.MinInt32 {
		return nil
	}

	node := TreeNode{
		Val: (*nums)[0],
	}
	newNums := (*nums)[1:]
	*nums = newNums
	node.Left = makeTree(nums)

	newNums = (*nums)[1:]
	*nums = newNums
	node.Right = makeTree(nums)

	return &node
}
