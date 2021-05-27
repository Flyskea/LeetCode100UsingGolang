/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 */
package main

import (
	"fmt"
	"math"
)

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

func maxPathSumCore(root *TreeNode, maxSum *int) int {
	// 如果当前节点为叶子节点，那么对父亲贡献为 0
	if root == nil {
		return 0
	}
	// 如果不是叶子节点，计算当前节点的左右孩子对自身的贡献left和right
	left := max(maxPathSumCore(root.Left, maxSum), 0)
	right := max(maxPathSumCore(root.Right, maxSum), 0)
	// 更新最大值，就是当前节点的val 加上左右节点的贡献。
	sum := root.Val + left + right
	if *maxSum < sum {
		*maxSum = sum
	}
	// 父节点不能左右都走不然会重复
	// 计算当前节点能为父亲提供的最大贡献，必须是把 val 加上！
	return root.Val + max(left, right)
}

func maxPathSum(root *TreeNode) int {
	max := math.MinInt64
	maxPathSumCore(root, &max)
	return max
}

// @lc code=end
func main() {
	fmt.Println(maxPathSum(nil))
}
