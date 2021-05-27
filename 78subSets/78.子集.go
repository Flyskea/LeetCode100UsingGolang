/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */
package main

import "fmt"

// @lc code=start
func backtracking(nums []int, k, start int, c []int, res *[][]int) {
	if len(c) == k {
		*res = append(*res, append([]int{}, c...))
		return
	}
	for i := start; i < len(nums); i++ {
		c = append(c, nums[i])
		backtracking(nums, k, i+1, c, res)
		c = c[:len(c)-1]
	}
}

func subsets(nums []int) [][]int {
	c, res := []int{}, [][]int{}
	for k := 0; k <= len(nums); k++ {
		backtracking(nums, k, 0, c, &res)
	}
	return res
}

// @lc code=end
func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
