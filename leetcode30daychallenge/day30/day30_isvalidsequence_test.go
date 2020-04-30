package day30

import (
	"math"
	"testing"
)

type testData struct {
	tree     []int
	input    []int
	expected bool
}

func TestIsValidSequence(t *testing.T) {
	m := math.MinInt32
	testDatas := []testData{
		{
			tree:     []int{0, 1, 0, 1, m, m, m, 1, 0, m, m, 0, m, m, 0, 0, m, m, m},
			input:    []int{0, 1, 0, 1},
			expected: true,
		},
		{
			tree:     []int{0, 1, 0, m, 1, m, m, 1, 0, m, m, 0, m, m, 0, 0, m, m, m},
			input:    []int{0, 0, 1},
			expected: false,
		},
		{
			tree:     []int{0, 1, 0, m, 1, m, m, 1, 0, m, m, 0, m, m, 0, 0, m, m, m},
			input:    []int{0, 1, 1},
			expected: false,
		},
		{
			tree:     []int{0, 1, 0, m, 1, m, m, 1, 0, m, m, 0, m, m, 0, 0, m, m, m},
			input:    []int{0, 1, 1},
			expected: false,
		},
	}

	for _, td := range testDatas {
		root := makeTree(&td.tree)
		actual := isValidSequence(root, td.input)

		if actual != td.expected {
			t.Errorf("Expected %t but was %t", td.expected, actual)
		}
	}
}

func makeTree(nums *[]int) *TreeNode {
	if len(*nums) <= 0 || (*nums)[0] == math.MinInt32 {
		return nil
	}

	node := TreeNode{
		Val: (*nums)[0],
	}

	if len(*nums) > 0 {
		newNums := (*nums)[1:]
		*nums = newNums
		node.Left = makeTree(nums)
	}

	if len(*nums) > 0 {
		newNums := (*nums)[1:]
		*nums = newNums
		node.Right = makeTree(nums)
	}

	return &node
}
