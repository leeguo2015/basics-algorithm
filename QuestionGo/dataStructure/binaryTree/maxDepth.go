package binaryTree

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	right := MaxDepth(root.Right)
	left := MaxDepth(root.Left)
	if right > left {
		return right + 1
	}
	return left + 1
}
