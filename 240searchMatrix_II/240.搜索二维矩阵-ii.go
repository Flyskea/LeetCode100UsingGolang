/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */
package main

import "fmt"

/*
func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			low, high := 0, len(matrix[0])-1
			if matrix[i][high] < target {
				break
			}
			for low <= high {
				mid := low + (high-low)>>1
				if matrix[i][mid] == target {
					return true
				} else if matrix[i][mid] < target {
					low = mid + 1
				} else {
					high = mid - 1
				}
			}
		}
	}
	return false
}
*/
// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	row, col := 0, cols-1
	for row < rows && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			row++
		} else {
			col--
		}
	}
	return false
}

// @lc code=end
func main() {
	fmt.Println(searchMatrix(nil, 1))
}
