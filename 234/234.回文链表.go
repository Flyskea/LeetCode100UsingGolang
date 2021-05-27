/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
递归
var front *ListNode

func isPalindromeCore(head *ListNode) bool {
	if head == nil {
		return true
	}
	if !isPalindromeCore(head.Next) {
		return false
	}
	if front.Val != head.Val {
		return false
	}
	front = front.Next
	return true
}

func isPalindrome(head *ListNode) bool {
	front = head
	return isPalindromeCore(head)
} */

// @lc code=start
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre, cur, next *ListNode
	cur = head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func findHalf(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	firstHalf := findHalf(head)
	secondHalf := reverseList(firstHalf.Next)
	p1, p2 := head, secondHalf
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	firstHalf.Next = reverseList(secondHalf)
	return true
}

// @lc code=end
func main() {
	fmt.Println(isPalindrome(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}))
}
