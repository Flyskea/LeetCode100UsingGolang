/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */
package main

import "fmt"

// @lc code=start
func canPartition(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum&1 == 1 {
		return false
	}
	target := sum / 2
	numsLen := len(nums)
	dp := make([][]bool, numsLen)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	for i := 1; i < numsLen; i++ {
		for j := 0; j <= target; j++ {
			dp[i][j] = dp[i-1][j]

			if nums[i] == j {
				dp[i][j] = true
				continue
			}
			if nums[i] < j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[numsLen-1][target]
}

// @lc code=end
func main() {
	fmt.Println(canPartition([]int{1, 2, 5}))
}
