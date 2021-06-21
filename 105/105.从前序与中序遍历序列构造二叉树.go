/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历结果
// [根节点, [左子树的前序遍历结果], [右子树的前序遍历结果]]
// 中序遍历结果
// [[左子树的中序遍历结果], 根节点, [右子树的中序遍历结果]]
// @lc code=start
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

// @lc code=end
func main() {
	fmt.Println(buildTree([]int{1}, []int{1}))
}
