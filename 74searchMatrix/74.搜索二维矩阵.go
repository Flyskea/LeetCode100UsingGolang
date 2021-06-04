/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */
package main

import "fmt"

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			low, high := 0, cols-1
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

// @lc code=end
func main() {
	fmt.Println(searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 23))
	fmt.Println(searchMatrix([][]int{{1, 1}}, 0))
}
