/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */
package main

import "fmt"

/*
DFS
var dx []int = []int{1, -1}

func dfs(nums []int, target, index int, ans *int) {
	if index == len(nums) {
		if target == 0 {
			*ans += 1
		}
		return
	}
	for i := 0; i < 2; i++ {
		dfs(nums, target-(dx[i]*nums[index]), index+1, ans)
	}
}
func findTargetSumWays(nums []int, target int) int {
	ans := 0
	dfs(nums, target, 0, &ans)
	return ans
} */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	diff := sum - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	n, neg := len(nums), diff/2
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1
	for i, num := range nums {
		for j := 0; j <= neg; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= num {
				dp[i+1][j] += dp[i][j-num]
			}
		}
	}
	return dp[n][neg]
}

// @lc code=end
func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 1))
}
