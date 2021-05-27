/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */
package main

import "fmt"

// @lc code=start
func nextPermutation(nums []int) {
	numsLen := len(nums)
	if numsLen < 2 {
		return
	}
	i, j, k := numsLen-2, numsLen-1, numsLen-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	for i, j := j, numsLen-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// @lc code=end
func main() {
	a := []int{1, 4, 3, 2}
	nextPermutation(a)
	fmt.Println(a)
}
