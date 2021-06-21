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

func main() {
	fmt.Println(GetUglyNumber_Solution(3))
	fmt.Println(dicesProbability(1))
}
