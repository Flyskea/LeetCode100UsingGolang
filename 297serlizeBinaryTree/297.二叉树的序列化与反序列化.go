/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
type Codec struct {
	l   []string
	str strings.Builder
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		c.str.WriteString("nil,")
	} else {
		c.str.WriteString(strconv.Itoa(root.Val))
		c.str.WriteString(",")
		c.serialize(root.Left)
		c.serialize(root.Right)
	}
	return c.str.String()
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	l := strings.Split(data, ",")
	for i := 0; i < len(l); i++ {
		if l[i] != "" {
			c.l = append(c.l, l[i])
		}
	}
	return c.rdeserialize()
}

func (c *Codec) rdeserialize() *TreeNode {
	if c.l[0] == "nil" {
		c.l = c.l[1:]
		return nil
	}

	val, _ := strconv.Atoi(c.l[0])
	root := &TreeNode{Val: val}
	c.l = c.l[1:]
	root.Left = c.rdeserialize()
	root.Right = c.rdeserialize()
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end
func main() {
	c := Constructor()
	fmt.Println(c.deserialize(c.serialize(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	})))
}
