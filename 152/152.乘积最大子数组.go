/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 考虑当前位置如果是一个负数的话
// 那么我们希望以它前一个位置结尾的某个段的积也是个负数，这样就可以负负得正，
// 并且我们希望这个积尽可能「负得更多」，即尽可能小
// 如果当前位置是一个正数的话
// 我们更希望以它前一个位置结尾的某个段的积也是个正数，并且希望它尽可能地大。
func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}
