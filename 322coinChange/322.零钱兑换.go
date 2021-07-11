/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */
package main

import "fmt"

// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	coinsLen := len(coins)
	for i := 1; i <= amount; i++ {
		for j := 0; j < coinsLen; j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// @lc code=end
func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 10))
}
