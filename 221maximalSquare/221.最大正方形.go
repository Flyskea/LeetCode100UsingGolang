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

// 之所以取最小的数，是因为要保证得到的是正方形而不是长方形呀。
// 比如从上往下有4个1，但是从左往右只有2个1，如果取4个1的话，
// 从左往右不够4个，只有两个，此时成了2*4的长方形。因此，只能取最小的那个2。
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
	for i := 0; i <= rows; i++ {
		for j := 0; j <= cols; j++ {
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
