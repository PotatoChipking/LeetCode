package _06

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	return reverse(inorder, postorder)
}

func reverse(inorder []int, postorder []int) *TreeNode {
	// 边界条件
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	// 创建当前子树的根节点
	// rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: postorder[len(postorder)-1]}

	// 在 inorder 中找到 rootVal 的位置
	inorderIndex := 0
	for inorderIndex < len(inorder) && inorder[inorderIndex] != postorder[len(postorder)-1] {
		inorderIndex++
	}

	// 划分左右子树，并递归构建
	// 中序：左根右
	// 后序：左右根
	// 中序遍历到根的位置，index在后序中代表左子树的数量，因此可直接拆分
	root.Left = reverse(inorder[:inorderIndex], postorder[:inorderIndex])
	root.Right = reverse(inorder[inorderIndex+1:], postorder[inorderIndex:len(postorder)-1])

	return root
}
