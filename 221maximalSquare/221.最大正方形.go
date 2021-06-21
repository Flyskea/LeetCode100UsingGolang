/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximalSquare(matrix [][]byte) int {
	rows, cols := len(matrix), len(matrix[0])
	if rows <= 0 || cols <= 0 {
		return 0
	}
	dp := make([][]int, rows+1)
	for i := range dp {
		dp[i] = make([]int, cols+1)
	}
	maxSide := 0
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
			maxSide = max(maxSide, dp[i][j])
		}
	}
	return maxSide * maxSide
}

// @lc code=end
func main() {
	fmt.Println(maximalSquare([][]byte{
		{'1', '1'},
		{'1', '1'},
	}))
}
