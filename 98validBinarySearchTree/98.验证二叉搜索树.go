/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 */
package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
直接判断左右值
func isValidBST(root *TreeNode) bool {
	return isValidbst(root, math.MinInt64, math.MaxInt64)
}
func isValidbst(root *TreeNode, min, max float64) bool {
	if root == nil {
		return true
	}
	v = root.Val
	return v < max && v > min && isValidbst(root.Left, min, v) && isValidbst(root.Right, v, max)
}
*/

// 中序遍历，数据一定是升序
// @lc code=start
func validBST(root *TreeNode, pre *int) bool {
	if root == nil {
		return true
	}
	if !validBST(root.Left, pre) {
		return false
	}
	if root.Val <= *pre {
		return false
	}
	*pre = root.Val
	return validBST(root.Right, pre)
}

func isValidBST(root *TreeNode) bool {
	pre := math.MinInt64
	return validBST(root, &pre)
}

// @lc code=end
