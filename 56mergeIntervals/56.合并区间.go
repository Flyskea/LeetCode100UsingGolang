/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */
package main

import (
	"fmt"
	"sort"
)

/*
先按左端点排序
第一个左端点是最小的
如果当前区间可以和merge数组最后一个合并，更新右端点
func merge(intervals [][]int) [][]int {
	var start, end int
	ans := [][]int{}
	intervalsLen := len(intervals)
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})
	for i := 0; i < intervalsLen; i++ {
		end = intervals[i][1]
		start = intervals[i][0]
		if len(ans) == 0 || ans[len(ans)-1][1] < start {
			ans = append(ans, []int{start, end})
		} else {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], end)
		}
	}
	return ans
} */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 双指针
func merge(intervals [][]int) [][]int {
	var j, end int
	ans := [][]int{}
	intervalsLen := len(intervals)
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})
	for i := 0; i < intervalsLen; {
		end = intervals[i][1]
		for j = i + 1; j < intervalsLen && intervals[j][0] <= end; j++ {
			end = max(end, intervals[j][1])
		}
		ans = append(ans, []int{intervals[i][0], end})
		i = j
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10},
		{15, 18}}))
}
