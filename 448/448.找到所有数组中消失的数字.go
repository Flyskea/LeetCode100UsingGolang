/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */
package main

import "fmt"

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	res := make([]int, -0)
	for _, v := range nums {
		if v < 0 {
			v = -v
		}
		if nums[v-1] > 0 {
			nums[v-1] = -nums[v-1]
		}
	}
	for i, v := range nums {
		if v > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

// @lc code=end
func main() {
	fmt.Println(findDisappearedNumbers([]int{1, 3, 3}))
}
