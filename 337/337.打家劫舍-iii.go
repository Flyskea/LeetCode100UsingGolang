/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
 */
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
// DFS
// 4 个孙子偷的钱 + 爷爷的钱 VS 两个儿子偷的钱 哪个组合钱多，就当做当前节点能偷的最大钱数。
// 这就是动态规划里面的最优子结构
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	money := root.Val
	if root.Left != nil {
		money += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		money += rob(root.Right.Left) + rob(root.Right.Right)
	}
	return max(money, rob(root.Left)+rob(root.Right))
}
*/

/*
// 记忆化
func robCore(root *TreeNode, memo map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if v, ok := memo[root]; ok {
		return v
	}
	money := root.Val
	if root.Left != nil {
		money += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		money += rob(root.Right.Left) + rob(root.Right.Right)
	}
	res := max(money, rob(root.Left)+rob(root.Right))
	memo[root] = res
	return res
}

func rob(root *TreeNode) int {
	memo := make(map[*TreeNode]int)
	return robCore(root, memo)
}
*/

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func robCore(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	left1, left2 := robCore(root.Left)
	right1, right2 := robCore(root.Right)
	// 不选当前节点
	noRoot := max(left1, left2) + max(right1, right2)
	// 选当前节点
	inRoot := root.Val + left1 + right1
	return noRoot, inRoot
}

func rob(root *TreeNode) int {
	a, b := robCore(root)
	return max(a, b)
}

// @lc code=end
func main() {
	fmt.Println(rob(nil))
}
