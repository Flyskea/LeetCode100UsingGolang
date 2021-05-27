/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */
package main

import "fmt"

// @lc code=start
func searchFirst(nums []int, target int) int {
	numsLen := len(nums)
	low, high := 0, numsLen
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] >= target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func searchLast(nums []int, target int) int {
	numsLen := len(nums)
	low, high := 0, numsLen
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low - 1
}

func searchRange(nums []int, target int) []int {
	numsLen := len(nums)
	if numsLen <= 0 {
		return []int{-1, -1}
	}
	low, high := searchFirst(nums, target), searchLast(nums, target)
	if low == numsLen || nums[low] != target {
		return []int{-1, -1}
	}
	return []int{low, high}
}

// @lc code=end
func main() {
	fmt.Println(searchRange([]int{1, 2, 3, 3, 4, 5}, 3))
}
