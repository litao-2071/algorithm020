/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// terminator
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	// dirll down
	leftRes := lowestCommonAncestor(root.Left, p, q)
	rightRes := lowestCommonAncestor(root.Right, p, q)
	// process in current level
	if leftRes != nil && rightRes != nil {
		return root
	}
	if leftRes != nil {
		return leftRes
	}
	return rightRes
}

// @lc code=end

