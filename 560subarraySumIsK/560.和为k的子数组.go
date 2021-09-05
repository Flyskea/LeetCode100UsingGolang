/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 */
package main

import "fmt"

// 前缀和：nums 的第 0 项到 当前项 的和。
// 用数组 prefixSum 表示，prefixSum[x]：第 0 项到 第 x 项 的和。
// prefixSum[x] = nums[0] + nums[1] + …+ nums[x]
// nums 的某项 = 两个相邻前缀和的差：
// nums[x] = prefixSum[x] - prefixSum[x - 1]
// nums 的 第 i 到 j 项 的和，有：
// nums[i] + … +nums[j] = prefixSum[j] - prefixSum[i - 1]
// 当 i 为 0，此时 i-1 为 -1，我们故意让 prefixSum[-1] 为 0，使得通式在i=0时也成立：
// nums[0] +…+nums[j] = prefixSum[j]
// @lc code=start
func subarraySum(nums []int, k int) int {
	count := 0
	hash := map[int]int{0: 1}
	preSum := 0

	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		if hash[preSum-k] > 0 {
			count += hash[preSum-k]
		}
		hash[preSum]++
	}
	return count
}

// @lc code=end
func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
}
