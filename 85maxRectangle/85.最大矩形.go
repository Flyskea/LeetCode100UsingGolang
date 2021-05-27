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
