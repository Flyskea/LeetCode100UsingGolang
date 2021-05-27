/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */
package main

import "fmt"

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canJump(nums []int) bool {
	rightmost := 0
	numsLen := len(nums)
	if numsLen < 2 {
		return true
	}
	for i, v := range nums {
		if i <= rightmost {
			rightmost = max(rightmost, i+v)
			if rightmost >= numsLen-1 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}

// @lc code=end
func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
