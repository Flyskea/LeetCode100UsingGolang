/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */
package main

import "fmt"

// @lc code=start
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

// @lc code=end
func main() {
	array := []int{1, 0, 3, 0, 4}
	moveZeroes(array)
	fmt.Println(array)
}
