/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func symmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil && right != nil {
		return false
	} else if left != nil && right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetric(root.Left, root.Right)
}

// @lc code=end
func main() {
	fmt.Println(isSymmetric(nil))
}
