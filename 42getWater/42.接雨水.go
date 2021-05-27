/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */
package main

import "fmt"

/*
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
*/

/* 按行求
func trap(height []int) int {
	var (
		max, heightLen, tmp, sum int
		flag                     bool
	)
	heightLen = len(height)
	for i := 0; i < heightLen; i++ {
		if height[i] > max {
			max = height[i]
		}
	}
	for i := 0; i <= max; i++ {
		tmp = 0
		// 已经找到第一个大于当前行的数
		flag = false
		for j := 0; j < heightLen; j++ {
			// 比当前行数小，可以积水
			if flag && height[j] < i {
				tmp++
			}
			if height[j] >= i {
				sum += tmp
				tmp, flag = 0, true
			}
		}
	}
	return sum
} */

/*
按列计算
只要左边最大和右边最大中最小的大于当前列，当前列就可积水
func trap(height []int) int {
	var sum, maxLeft, maxRight, minHeight int
	heightLen := len(height)
	for i := 1; i < heightLen-1; i++ {
		maxLeft, maxRight = 0, 0
		for j := i - 1; j >= 0; j-- {
			if height[j] > maxLeft {
				maxLeft = height[j]
			}
		}
		for j := i + 1; j < heightLen; j++ {
			if height[j] > maxRight {
				maxRight = height[j]
			}
		}
		minHeight = min(maxLeft, maxRight)
		if height[i] < minHeight {
			sum += minHeight - height[i]
		}
	}
	return sum
} */

/* // 一次遍历，保存第i列的左边最高和右边最高，优化按列计算复杂度
func trap(height []int) (ans int) {
	n := len(height)
	if n == 0 {
		return
	}

	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	for i, h := range height {
		ans += min(leftMax[i], rightMax[i]) - h
	}
	return
} */

// @lc code=start

// 对于位置left而言，它左边最大值一定是left_max，右边最大值“大于等于”right_max，
// 这时候，如果left_max<right_max成立，
// 那么它就知道自己能存多少水了。
// 无论右边将来会不会出现更大的right_max，
// 都不影响这个结果。
// 所以当left_max<right_max时，
// 我们就希望去处理left下标，
// 反之，我们希望去处理right下标。
func trap(height []int) int {
	res, left, right, maxLeft, maxRight := 0, 0, len(height)-1, 0, 0
	for left <= right {
		if height[left] <= height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				res += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				res += maxRight - height[right]
			}
			right--
		}
	}
	return res
}

// @lc code=end

func main() {
	fmt.Println(trap([]int{2, 1, 0, 1, 2}))
}
