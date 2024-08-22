package binary_tree_maximum_path_sum

import "math"

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

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt64

	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	reverse(root, &ans)
	return ans
}

func reverse(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}

	l := max(reverse(root.Left, ans), 0)
	r := max(reverse(root.Right, ans), 0)
	cur := max(l, r)
	sum := max(root.Val, root.Val+l, root.Val+r, root.Val+l+r)
	if sum > *ans {
		*ans = sum
	}
	return root.Val + cur
}
