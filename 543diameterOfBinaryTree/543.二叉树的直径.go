/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func diameterOfBinaryTreeCore(root *TreeNode, maxDepth *int) int {
	if root == nil {
		return 0
	}
	left := diameterOfBinaryTreeCore(root.Left, maxDepth)
	right := diameterOfBinaryTreeCore(root.Right, maxDepth)
	if *maxDepth < left+right {
		*maxDepth = left + right
	}
	return max(left, right) + 1
}
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDepth := 0
	diameterOfBinaryTreeCore(root, &maxDepth)
	return maxDepth
}

// @lc code=end
func main() {
	fmt.Println(diameterOfBinaryTree(nil))
}
