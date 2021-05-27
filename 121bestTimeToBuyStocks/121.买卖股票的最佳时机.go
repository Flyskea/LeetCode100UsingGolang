/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
package main

import "fmt"

// @lc code=start
func maxProfit(prices []int) int {
	min := prices[0]
	max := 0
	for i := 1; i < len(prices); i++ {
		if max < prices[i]-min {
			max = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return max
}

// @lc code=end
func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
