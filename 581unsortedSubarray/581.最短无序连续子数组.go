/*
 * @lc app=leetcode.cn id=581 lang=golang
 *
 * [581] 最短无序连续子数组
 */
package main

import (
	"fmt"
	"math"
)

/*
暴力寻找左边界与右边界
左边界为需要被交换的数字的最小下标
右边界为需要被交换的数字的最大下标
func findUnsortedSubarray(nums []int) int {
	var (
		start, end int
	)
	numsLen := len(nums)
	start = numsLen
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < numsLen; j++ {
			if nums[j] < nums[i] {
				start = min(start, i)
				end = max(end, j)
			}
		}
	}
	if end-start > 0 {
		return end - start + 1
	}
	return 0
}
*/

/*
// 栈
	var (
		start, end int
	)
	numsLen := len(nums)
	stack := []int{}
	start = numsLen
	for i := 0; i < len(nums); i++ {
		for len(stack) != 0 && nums[stack[len(stack)-1]] > nums[i] {
			start = min(start, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = []int{}
	for j := numsLen - 1; j >= 0; j-- {
		for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[j] {
			end = max(end, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, j)
	}
	if end-start > 0 {
		return end - start + 1
	}
	return 0
*/
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

// 不使用额外空间
// 先从左边找到第一个降序的元素，并记录最小的元素 min，
// 再从右边往左找到最右边开始降序的元素，并记录最大的元素 max。
// 最后需要还原最小元素和最大元素在原数组中正确的位置。
// 以逆序区间左边界为例，如果区间外的一个元素比这个逆序区间内的最小元素还要小，说明它并不是左边界，
// 因为这个小元素和逆序区间内的最小元素组合在一起，还是升序，并不是逆序。
// 只有在左边区间外找到第一个大于逆序区间内最小元素，说明这里刚刚开始发生逆序，这才是最小逆序区间的左边界。
// 同理，在逆序区间的右边找到第一个小于逆序区间内最大元素，说明这里刚刚发生逆序，这才是最小逆序区间的右边界。
// 至此，最小逆序区间的左右边界都确定下来了，最短长度也就确定了下来。时间复杂度 O(n)，空间复杂度 O(1)。
func findUnsortedSubarray(nums []int) int {
	n, left, right, minR, maxL := len(nums), -1, -1, math.MaxInt32, math.MinInt32
	// left
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			minR = min(minR, nums[i])
			maxL = max(maxL, nums[i-1])
		}
	}
	// minR
	for i := 0; i < n; i++ {
		if nums[i] > minR {
			left = i
			break
		}
	}
	// maxL
	for i := n - 1; i >= 0; i-- {
		if nums[i] < maxL {
			right = i
			break
		}
	}
	if left == -1 || right == -1 {
		return 0
	}
	return right - left + 1
}

// @lc code=end
func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 1}))
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4}))
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 3, 3}))
}
