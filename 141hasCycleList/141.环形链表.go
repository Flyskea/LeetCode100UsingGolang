/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head.Next
	if slow == nil {
		return false
	}
	fast := slow.Next
	for fast != nil && slow != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
		slow = slow.Next
	}
	return false
}

// @lc code=end
func main() {
	fmt.Println(hasCycle(nil))
}
