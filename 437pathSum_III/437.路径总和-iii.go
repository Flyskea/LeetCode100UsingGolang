/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
DFS
func pathSumCore(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	ans := 0
	targetSum -= root.Val
	if targetSum == 0 {
		ans += 1
	}
	ans += pathSumCore(root.Left, targetSum)
	ans += pathSumCore(root.Right, targetSum)
	return ans
}
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return pathSum(root.Left, targetSum) +
		pathSum(root.Right, targetSum) +
		pathSumCore(root, targetSum)
}
*/

// 路径方向必须是向下的（只能从父节点到子节点）
// 当我们讨论两个节点的前缀和差值时，有一个前提：
// 一个节点必须是另一个节点的祖先节点
// 前缀和+DFS
// @lc code=start
func dfs(root *TreeNode, prefixSum map[int]int, cur, sum int) int {
	if root == nil {
		return 0
	}
	cur += root.Val
	cnt := 0
	if v, ok := prefixSum[cur-sum]; ok {
		cnt = v
	}
	prefixSum[cur]++
	cnt += dfs(root.Left, prefixSum, cur, sum)
	cnt += dfs(root.Right, prefixSum, cur, sum)
	prefixSum[cur]--
	return cnt
}

func pathSum(root *TreeNode, targetSum int) int {
	prefixSum := make(map[int]int)
	prefixSum[0] = 1
	return dfs(root, prefixSum, 0, targetSum)
}

// @lc code=end
func main() {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: -2,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: -3,
			Right: &TreeNode{
				Val: 11,
			},
		},
	}
	fmt.Println(pathSum(root, 8))
}
