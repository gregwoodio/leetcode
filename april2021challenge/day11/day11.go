package day11

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	sums := []int{}
	depth := 0

	sums = *(traverse(root, depth, &sums))

	return sums[len(sums)-1]
}

func traverse(node *TreeNode, depth int, sums *[]int) *[]int {
	if node == nil {
		return sums
	}

	// check if we need to add to our sums slice
	if depth == len(*sums) {
		newSums := append(*sums, 0)
		sums = &newSums
	}

	(*sums)[depth] += node.Val

	sums = traverse(node.Left, depth+1, sums)
	sums = traverse(node.Right, depth+1, sums)

	return sums
}
