/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */
package main

import "fmt"

// @lc code=start
func singleNumber(nums []int) int {
	var ans int
	for _, v := range nums {
		ans = v ^ ans
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(singleNumber([]int{2, 2, 1, 1, 4}))
}
