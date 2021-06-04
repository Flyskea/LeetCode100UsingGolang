/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */
package main

import "fmt"

/*
// 快慢指正，想象成一个链表
// 数组中数字代表下一个结点的数组下标
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	slow = nums[slow]
	fast = nums[nums[fast]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	p1, p2 := 0, slow
	for p1 != p2 {
		p1 = nums[p1]
		p2 = nums[p2]
	}
	return p1
}
*/
// 二分搜索
// 假设有 n+1 个数，则可能重复的数位于区间 [1, n] 中。记该区间最小值、最大值和中间值为 low、high、mid
// 遍历整个数组，统计小于等于 mid 的整数的个数，至多为 mid 个
// 如果超过 mid 个就说明重复的数存在于区间 [low,mid] （闭区间）中；否则，重复的数存在于区间 (mid, high] （左开右闭）中
// 缩小区间，继续重复步骤 2、3，直到区间变成 1 个整数，即 low == high
// 整数 low 就是要找的重复的数
// @lc code=start
func findDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid, count := low+(high-low)>>1, 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		if count > mid {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

// @lc code=end
func main() {
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
}
