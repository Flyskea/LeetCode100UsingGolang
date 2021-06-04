/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */
package main

import "fmt"

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	numsLen := len(nums)
	if numsLen < 2 {
		return nums
	}
	// 存数字下标
	queue := make([]int, 0, k)
	ans := make([]int, 0, numsLen-k+1)
	for i, v := range nums {
		// 队列头超出窗口大小，移除
		if i >= k && queue[0] == i-k {
			queue = queue[1:]
		}
		// 队列按大到小排序
		for len(queue) != 0 && nums[queue[len(queue)-1]] < v {
			queue = queue[0 : len(queue)-1]
		}
		queue = append(queue, i)
		// 队列头即为答案
		if i >= k-1 {
			ans = append(ans, nums[queue[0]])
		}
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
