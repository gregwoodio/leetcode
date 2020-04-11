package day11

import (
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	nums := []int{1, 2, 4, -1, -1, 5, -1, -1, 3, -1, -1}
	root := makeTree(&nums)

	actual := diameterOfBinaryTree(root)
	expected := 3

	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func makeTree(nums *[]int) *TreeNode {
	if len(*nums) <= 0 || (*nums)[0] == -1 {
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
