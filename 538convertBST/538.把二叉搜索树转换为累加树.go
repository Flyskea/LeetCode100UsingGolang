/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func convertBSTCore(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	convertBSTCore(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	convertBSTCore(root.Left, sum)
}
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	convertBSTCore(root, &sum)
	return root
}

// @lc code=end
func main() {
	fmt.Println(convertBST(nil))
}
