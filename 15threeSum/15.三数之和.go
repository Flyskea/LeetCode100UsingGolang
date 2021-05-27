/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
package main

import (
	"fmt"
	"sort"
)

// @lc code=start
func threeSum(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	numsLen := len(nums)
	if numsLen < 3 {
		return nil
	}
	ans := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < numsLen-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := numsLen - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(threeSum([]int{-2, -2, -2, -1, -1, 2, 4, 4}))
}
