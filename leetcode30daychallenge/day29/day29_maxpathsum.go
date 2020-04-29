package day29

import "math"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	sum := math.MinInt32
	helper(root, &sum)
	return sum
}

func helper(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}

	left, right := 0, 0
	if node.Left != nil {
		left = helper(node.Left, maxSum)
	}
	if node.Right != nil {
		right = helper(node.Right, maxSum)
	}

	leftOrRightOrNode := max(node.Val, node.Val+max(left, right))
	*maxSum = max(*maxSum, max(node.Val+left+right, leftOrRightOrNode))
	return leftOrRightOrNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
