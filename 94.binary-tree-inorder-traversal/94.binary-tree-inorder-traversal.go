/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var inorderNodes []int
	var leftInorder []int
	var rightInorder []int

	if root == nil {
		return inorderNodes
	}

	if root.Left == nil {
		leftInorder = []int{}
	} else {
		leftInorder = inorderTraversal(root.Left)
	}

	if root.Right == nil {
		rightInorder = []int{}
	} else {
		rightInorder = inorderTraversal(root.Right)
	}

	inorderNodes = append(inorderNodes, leftInorder...)
	inorderNodes = append(inorderNodes, root.Val)
	inorderNodes = append(inorderNodes, rightInorder...)

	return inorderNodes
}

// @lc code=end

