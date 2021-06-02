/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func pathSumHelper(root *TreeNode, targetSum int, ans [][]int, stack []int) [][]int {
	if root == nil {
		return ans
	}
	targetSum -= root.Val
	stack = append(stack, root.Val)
	if targetSum == 0 && root.Left == nil && root.Right == nil {
		ans = append(ans, append([]int{}, stack...))
	}
	ans = pathSumHelper(root.Left, targetSum, ans, stack)
	ans = pathSumHelper(root.Right, targetSum, ans, stack)
	return ans
}
func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := [][]int{}
	ans = pathSumHelper(root, targetSum, ans, []int{})
	return ans
}

// @lc code=end
func main() {
	fmt.Println(pathSum(nil, 0))
}
