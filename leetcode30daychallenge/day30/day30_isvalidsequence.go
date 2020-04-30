package day30

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidSequence(root *TreeNode, arr []int) bool {

	if len(arr) == 0 {
		return false
	}

	node := root
	if node.Val != arr[0] {
		return false
	}

	return checkNode(node, arr, 1)
}

func checkNode(node *TreeNode, arr []int, i int) bool {
	if i == len(arr) {
		return node.Left == nil && node.Right == nil
	}

	leftPath, rightPath := false, false

	if node.Left != nil && node.Left.Val == arr[i] {
		leftPath = checkNode(node.Left, arr, i+1)
	}

	if node.Right != nil && node.Right.Val == arr[i] {
		rightPath = checkNode(node.Right, arr, i+1)
	}

	return leftPath || rightPath
}
