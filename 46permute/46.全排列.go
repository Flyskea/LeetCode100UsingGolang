/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */
package main

import (
	"fmt"
)

// @lc code=start
func backtracking(nums, res []int, ans *[][]int, used *[]bool) {
	numsLen := len(nums)
	if len(res) == numsLen {
		*ans = append(*ans, append([]int{}, res...))
	}
	for i := 0; i < numsLen; i++ {
		if !(*used)[i] {
			(*used)[i] = true
			res = append(res, nums[i])
			backtracking(nums, res, ans, used)
			res = res[:len(res)-1]
			(*used)[i] = false
		}
	}
}

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	used := make([]bool, len(nums))
	backtracking(nums, []int{}, &ans, &used)
	return ans
}

// @lc code=end
func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
