/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */
package main

import (
	"fmt"
)

/*
DFS
var dx []int = []int{0, 1}
var dy []int = []int{1, 0}

func dfs(grid [][]int, x, y, sum int, minSum *int) {
	rows, cols := len(grid), len(grid[0])
	if x == rows-1 && y == cols-1 {
		if sum < *minSum {
			*minSum = sum
		}
		return
	}
	for i := 0; i < 2; i++ {
		row, col := x+dx[i], y+dy[i]
		if row >= rows || col >= cols {
			continue
		}
		dfs(grid, row, col, sum+grid[row][col], minSum)
	}
}
func minPathSum(grid [][]int) int {
	min := math.MaxInt64
	dfs(grid, 0, 0, grid[0][0], &min)
	return min
} */

// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func minPathSum(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[rows-1][cols-1]
}

// @lc code=end
func main() {
	fmt.Println(minPathSum([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))
}
