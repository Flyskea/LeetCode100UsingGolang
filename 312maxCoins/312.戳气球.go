/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
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
func maxCoins(nums []int) int {
	numsLen := len(nums)
	rec := make([][]int, numsLen+2)
	for i := 0; i < numsLen+2; i++ {
		rec[i] = make([]int, numsLen+2)
	}
	val := make([]int, numsLen+2)
	val[0], val[numsLen+1] = 1, 1
	for i := 1; i <= numsLen; i++ {
		val[i] = nums[i-1]
	}
	for i := numsLen - 1; i >= 0; i-- {
		for j := i + 2; j <= numsLen+1; j++ {
			for k := i + 1; k < j; k++ {
				sum := val[i] * val[k] * val[j]
				sum += rec[i][k] + rec[k][j]
				rec[i][j] = max(rec[i][j], sum)
			}
		}
	}
	return rec[0][numsLen+1]
}

// @lc code=end
func main() {
	fmt.Println(maxCoins([]int{1, 5}))
}
