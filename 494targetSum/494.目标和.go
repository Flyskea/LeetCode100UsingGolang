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

/*
原问题等同于： 找到nums一个正子集P和一个负子集N，使得总和等于target。即sum(P) - sum(N) == target，
即sum(P) + sum(N) + sum(P) - sum(N) == target + sum(P) + sum(N)
即2 * sum(P) == target + sum(nums)， 其中target + sum(nums)必须>=0且为偶数，否则等式不可能成立。
则问题转换为：存在多少个子集P，使sum(P) == (target + sum(nums))/2。

dp[i][j]表示前i个元素有多少个目标和为j的子集。dp[0][0] = 1
    1. dp[i][j] = dp[i-1][j]
    2. 如果nums[0...i-2]存在目标和为j-nums[i-1]的子集，则dp[i][j] += dp[i-1][j-nums[i-1]]

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
} */

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	diff := sum - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	neg := diff / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	for _, num := range nums {
		for j := neg; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[neg]
}

// @lc code=end
func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
