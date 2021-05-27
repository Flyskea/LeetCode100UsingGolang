/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */
package main

import (
	"fmt"
)

// @lc code=start
func rob(nums []int) int {
	numsLen := len(nums)
	dp := make([]int, numsLen)
	if numsLen == 0 {
		return 0
	}
	if numsLen == 1 {
		return nums[0]
	}
	dp[0] = nums[0]
	if nums[0] > nums[1] {
		dp[1] = nums[0]
	} else {
		dp[1] = nums[1]
	}
	for i := 2; i < numsLen; i++ {
		if dp[i-1] > (nums[i] + dp[i-2]) {
			dp[i] = dp[i-1]
		} else {
			dp[i] = nums[i] + dp[i-2]
		}
	}
	return dp[numsLen-1]
}

// @lc code=end
func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
}
