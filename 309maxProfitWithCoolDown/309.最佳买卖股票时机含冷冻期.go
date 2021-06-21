/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
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
func maxProfit(prices []int) int {
	pricesLen := len(prices)
	dp := make([][3]int, pricesLen+1)
	dp[0][0] = -prices[0]
	for i := 1; i < pricesLen; i++ {
		// f[i][0]: 手上持有股票的最大收益
		// f[i][1]: 手上不持有股票，并且处于冷冻期中的累计最大收益
		// f[i][2]: 手上不持有股票，并且不在冷冻期中的累计最大收益
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	ans := max(dp[pricesLen-1][1], dp[pricesLen-1][2])
	return ans
}

// @lc code=end
func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
}
