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
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	return selection(nums, 0, len(nums)-1, len(nums)-k)
}

func selection(nums []int, l, r, k int) int {
	if l == r {
		return nums[l]
	}
	p := partition(nums, l, r)
	if k == p {
		return nums[p]
	} else if k < p {
		return selection(nums, l, p-1, k)
	} else {
		return selection(nums, p+1, r, k)
	}

}

func partition(a []int, low, high int) int {
	pivot := a[high]
	i := low - 1
	for j := low; j < high; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[high] = a[high], a[i+1]
	return i + 1
}

// @lc code=end
func main() {
	fmt.Println(findKthLargest([]int{1, 5, 3}, 2))
}
