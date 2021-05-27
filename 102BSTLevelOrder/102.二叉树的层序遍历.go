/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{}
	ans := [][]int{}
	queue = append(queue, root)
	level := 0
	for len(queue) != 0 {
		ans = append(ans, []int{})
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			v := queue[0]
			queue = queue[1:]
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
			ans[level] = append(ans[level], v.Val)
		}
		level++
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(levelOrder(nil))
}
