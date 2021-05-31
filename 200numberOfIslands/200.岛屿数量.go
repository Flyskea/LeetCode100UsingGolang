/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */
package main

import "fmt"

// @lc code=start
var dx []int = []int{0, -1, 1, 0}
var dy []int = []int{1, 0, 0, -1}

func dfs(grid [][]byte, row, col int) {
	if grid[row][col] == '1' {
		grid[row][col] = '0'
		var x, y int
		for i := 0; i < 4; i++ {
			x, y = row+dx[i], col+dy[i]
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == '1' {
				dfs(grid, x, y)
			}
		}
	}
}
func numIslands(grid [][]byte) int {
	var res int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

// @lc code=end
func main() {
	fmt.Println(numIslands([][]byte{{'1', '1', '1', '1', '0'}}))
}
