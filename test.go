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

type ListNode struct {
	Val  int
	Next *ListNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dicesProbability(n int) []float64 {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 6*n+1)
	}
	for i := 1; i < 7; i++ {
		dp[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := i; j <= 6*i; j++ {
			for cur := 1; cur < 7 && cur <= j; cur++ {
				dp[i][j] += dp[i-1][j-cur]
			}
		}
	}
	all := math.Pow(6, float64(n))
	res := make([]float64, 0)
	for i := n; i <= 6*n; i++ {
		res = append(res, float64(dp[n][i])/all)
	}
	return res
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

func GetUglyNumber_Solution(index int) int {
	if index <= 1 {
		return index
	}
	// write code here
	uglyNumbers, a, b, c := make([]int, index), 0, 0, 0
	for i := range uglyNumbers {
		uglyNumbers[i] = 1
	}
	for i := 1; i < index; i++ {
		n2, n3, n5 := uglyNumbers[a]*2, uglyNumbers[b]*3, uglyNumbers[c]*5
		uglyNumbers[i] = min3(n2, n3, n5)
		if uglyNumbers[i] == n2 {
			a++
		}
		if uglyNumbers[i] == n3 {
			b++
		}
		if uglyNumbers[i] == n5 {
			c++
		}
	}
	return uglyNumbers[len(uglyNumbers)-1]
}

func min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}

func maxProduct(nums []int) int {
	numsLen := len(nums)
	dp := make([][2]int, numsLen)
	dp[0][0] = nums[0]
	dp[0][1] = nums[0]
	ans := 0
	for i := 1; i < numsLen; i++ {
		dp[i][0] = max(dp[i-1][0]*nums[i], max(nums[i], nums[i]*dp[i-1][1]))
		dp[i][1] = min(dp[i-1][1]*nums[i], min(nums[i], nums[i]*dp[i-1][0]))
		ans = max(ans, dp[i][0])
	}
	return ans
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Printf("%d", root.Val)
	inorder(root.Right)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	return root
}

type doubleLinkList struct {
	key, val   int
	prev, next *doubleLinkList
}
type LRUCache struct {
	head, tail *doubleLinkList
	keys       map[int]*doubleLinkList
	capacity   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity,
		keys: make(map[int]*doubleLinkList),
	}
}

func (lru *LRUCache) add(node *doubleLinkList) {
	node.prev = nil
	node.next = lru.head
	if lru.head != nil {
		lru.head.prev = node
	}
	lru.head = node
	if lru.tail == nil {
		lru.tail = node
		lru.tail.next = nil
	}
}

func (lru *LRUCache) remove(node *doubleLinkList) {
	if node == lru.head {
		lru.head = node.next
		if node.next != nil {
			node.next.prev = nil
		}
		node.next = nil
		return
	}
	if node == lru.tail {
		lru.tail = node.prev
		node.prev.next = nil
		node.prev = nil
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.keys[key]; ok {
		lru.remove(node)
		lru.add(node)
		return node.val
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.keys[key]; ok {
		lru.remove(node)
		lru.add(node)
		node.val = value
	} else {
		node = &doubleLinkList{
			val: value,
			key: key,
		}
		lru.keys[key] = node
		lru.add(node)
	}
	if len(lru.keys) > lru.capacity {
		delete(lru.keys, lru.tail.key)
		lru.remove(lru.tail)
	}
}

func main() {
	a := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	inorder(a)
	fmt.Println()
	invertTree(a)
	inorder(a)
	lru := Constructor(3)
	lru.Get(2)
	lru.Put(2, 2)
	lru.Get(2)
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}
