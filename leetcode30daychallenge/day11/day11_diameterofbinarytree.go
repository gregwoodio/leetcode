package day11

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	leftDiam := diameterOfBinaryTree(root.Left)
	rightDiam := diameterOfBinaryTree(root.Right)

	return max(leftHeight+rightHeight, max(leftDiam, rightDiam))
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := height(node.Left)
	rightHeight := height(node.Right)

	return 1 + max(leftHeight, rightHeight)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
