/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
func lowestCommonAncestorCore(root, p, q *TreeNode, ancestor **TreeNode) {
	if root == nil {
		return
	}
	if p.Val < root.Val && q.Val < root.Val {
		lowestCommonAncestorCore(root.Left, p, q, ancestor)
	} else if p.Val > root.Val && q.Val > root.Val {
		lowestCommonAncestorCore(root.Right, p, q, ancestor)
	} else {
		*ancestor = root
	}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ancestor := &TreeNode{}
	lowestCommonAncestorCore(root, p, q, &ancestor)
	return ancestor
} */

// @lc code=start
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ancestor := root
	for {
		if p.Val < ancestor.Val && q.Val < ancestor.Val {
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			break
		}
	}
	return ancestor
}

// @lc code=end
func main() {
	lowestCommonAncestor(nil, nil, nil)
}
