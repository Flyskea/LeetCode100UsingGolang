/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */
package main

import "fmt"

// @lc code=start
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}

// @lc code=end
func main() {
	matrix := [][]int{{1, 2, 3, 4}, {4, 5, 6, 7}, {7, 8, 9, 10}, {15, 16, 17, 18}}
	rotate(matrix)
	fmt.Println(matrix)
}
