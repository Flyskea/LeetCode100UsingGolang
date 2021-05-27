/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */
package main

import "fmt"

// @lc code=start
var dx []int = []int{-1, 0, 0, 1}
var dy []int = []int{0, -1, 1, 0}

func backtracking(board [][]byte, word string, row, col, index int, visited [][]bool) bool {
	if len(word)-1 == index {
		return board[row][col] == word[index]
	}
	var x, y, cols, rows int
	rows, cols = len(board), len(board[0])
	if board[row][col] == word[index] {
		visited[row][col] = true
		for i := 0; i < 4; i++ {
			x = row + dx[i]
			y = col + dy[i]
			if x >= 0 && x < rows && y >= 0 && y < cols && !visited[x][y] && backtracking(board, word, x, y, index+1, visited) {
				return true
			}
		}
		visited[row][col] = false
	}
	return false
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	for i, v := range board {
		for j := range v {
			if backtracking(board, word, i, j, 0, visited) {
				return true
			}
		}
	}
	return false
}

// @lc code=end
func main() {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCESEEDASFC"))
}
