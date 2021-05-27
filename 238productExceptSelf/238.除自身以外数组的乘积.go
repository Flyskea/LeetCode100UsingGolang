/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */
package main

import "fmt"

// @lc code=start
func productExceptSelf(nums []int) []int {
	numsLen := len(nums)
	ans := make([]int, numsLen)
	ans[0] = 1
	for i := 1; i < numsLen; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	r := 1
	for i := numsLen - 1; i >= 0; i-- {
		ans[i] *= r
		r *= nums[i]
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
