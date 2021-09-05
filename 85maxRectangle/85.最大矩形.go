/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */
package main

import (
	"fmt"
)

// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
将每一列高度统计，将参数传递给84题
func largestRectangleArea(heights []int) int {
	var maxArea, left, right, cur, curHeight, width int
	stack := []int{}
	for i, v := range heights {
		for len(stack) > 0 && stack[len(stack)-1] < v {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			right = i
			left = stack[len(stack)-1]
			width = right - left - 1
			curHeight = heights[cur]
			maxArea = max(maxArea, curHeight*width)
		}
		stack = append(stack, i)
	}
	return maxArea
}

func maximalRectangle1(martix [][]byte) int {
	heights := make([]int, len(martix))
	var maxArea int
	for i := range martix {
		for j, col := range martix[i] {
			if col == '1' {
				heights[j] += 1
			}
		}
		maxArea = max(maxArea, largestRectangleArea(heights))
	}
	return maxArea
} */

func maximalRectangle(matrix [][]byte) int {
	var ans int
	if len(matrix) == 0 {
		return 0
	}
	rows, cols := len(matrix), len(matrix[0])
	// 获取当前行当前列左边最多连续的1
	left := make([][]int, rows)
	for i, row := range matrix {
		left[i] = make([]int, cols)
		for j, col := range row {
			if col == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
			width := left[i][j]
			area := width
			// 向上累加 如果当前行列没有1直接跳出
			for k := i - 1; k >= 0; k-- {
				if left[k][j] == 0 {
					break
				}
				width = min(width, left[k][j])
				area = max(area, (i-k+1)*width)
			}
			ans = max(ans, area)
		}
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(maximalRectangle([][]byte{{'0'}}))
}
