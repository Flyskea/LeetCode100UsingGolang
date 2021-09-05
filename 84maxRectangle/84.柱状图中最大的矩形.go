/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
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

/*
单调栈
func largestRectangleArea(heights []int) int {
	var cur,curHeight, left, right, curWidth, ans int
	newHeights := make([]int, len(heights)+2)
	stack := make([]int, 0)
	for i := 1; i < len(heights)+1; i++ {
		newHeights[i] = heights[i-1]
	}
	for i, v := range newHeights {
		for len(stack) != 0 && v < newHeights[stack[len(stack)-1]] {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			curHeight = newHeights[cur]
			left = stack[len(stack)-1]
			right = i
			curWidth = right - left - 1
			ans = max(ans, curHeight*curWidth)
		}
		stack = append(stack, i)
	}
	return ans
} */

func largestRectangleArea(heights []int) int {
	var ans int
	stack := []int{}
	heightLen := len(heights)
	left, right := make([]int, heightLen), make([]int, heightLen)
	for i := 0; i < heightLen; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = []int{}
	for i := heightLen - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = heightLen
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i := 0; i < heightLen; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(largestRectangleArea([]int{0, 9}))
}
