package day20

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
	var head, node, prev *TreeNode

	for _, val := range preorder {
		node = newNode(val)

		if head == nil {
			head = node
			continue
		}

		prev = head

		for {
			if prev.Val > node.Val && prev.Left == nil {
				prev.Left = node
				break
			} else if prev.Val > node.Val {
				prev = prev.Left
			} else if prev.Val < node.Val && prev.Right == nil {
				prev.Right = node
				break
			} else if prev.Val < node.Val {
				prev = prev.Right
			}
		}
	}

	return head
}

func newNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}
