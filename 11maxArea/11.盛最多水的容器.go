/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */
package main

import "fmt"

// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxArea(height []int) int {
	heightLen := len(height)
	res := 0
	area := 0
	for i, j := 0, heightLen-1; i != j; {
		left, right := height[i], height[j]
		area = (j - i) * min(left, right)
		if left < right {
			i++
		} else {
			j--
		}
		if area > res {
			res = area
		}
	}
	return res
}

// @lc code=end

func main() {
	fmt.Println(maxArea([]int{1, 3, 2, 5, 25, 24, 5}))
}
