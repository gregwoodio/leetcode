package day11

import "testing"

func TestDeepestLeavesSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 2, 3, 4, 5, -1, 6, 7, -1, -1, -1, -1, 8},
			expected: 15,
		},
		{
			nums:     []int{6, 7, 8, 2, 7, 1, 3, 9, -1, 1, 4, -1, -1, -1, 5},
			expected: 19,
		},
	}

	for _, tc := range testCases {
		root := makeTree(tc.nums)
		actual := deepestLeavesSum(root)

		if actual != tc.expected {
			t.Errorf("Expected %d but was %d", tc.expected, actual)
		}
	}
}

// Makes a tree given Leetcode's format for trees:
// 1, 2, 3, 4, 5, -1, 6 would give this tree:
//     1
//    / \
//   2   3
//  / \   \
// 4   5   6
//
// To do this we make an array of all the nodes based on their
// index in the nums list. If nodes need to be processed still
// they are added to a queue.
func makeTree(nums []int) *TreeNode {

	nodes := []*TreeNode{}

	// make nodes same length as nums
	for range nums {
		nodes = append(nodes, nil)
	}

	queue := []int{}
	index := 0

	// populate our queue
	queue = append(queue, index)

	// create the first node. Others will be added in the loop
	nodes[index] = &TreeNode{
		Val: nums[index],
	}

	for {
		next := queue[0]
		queue = queue[1:]

		node := nodes[next]

		index++
		if index >= len(nums) {
			break
		}

		if nums[index] != -1 {
			queue = append(queue, index)

			nodes[index] = &TreeNode{
				Val: nums[index],
			}

			node.Left = nodes[index]
		}

		index++
		if index >= len(nums) {
			break
		}

		if nums[index] != -1 {
			queue = append(queue, index)

			nodes[index] = &TreeNode{
				Val: nums[index],
			}

			node.Right = nodes[index]
		}
	}

	return nodes[0]
}
