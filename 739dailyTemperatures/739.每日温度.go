/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */
package main

import (
	"fmt"
)

/*
func dailyTemperatures(temperatures []int) []int {
	var index int
	length := len(temperatures)
	ans := make([]int, length)
	// 保存温度第一次出现的位置
	next := make([]int, 101)
	for i := 0; i < 101; i++ {
		next[i] = math.MaxInt32
	}
	for i := length - 1; i >= 0; i-- {
		index = math.MaxInt32
		// 查看有没有大于当前天数温度的
		for t := temperatures[i] + 1; t <= 100; t++ {
			if next[t] < index {
				index = next[t]
			}
		}
		if index < math.MaxInt32 {
			ans[i] = index - i
		}
		next[temperatures[i]] = i
	}
	return ans
}
*/

// @lc code=start
// 单调栈
func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	var toCheck []int
	for i, t := range T {
		for len(toCheck) > 0 && T[toCheck[len(toCheck)-1]] < t {
			idx := toCheck[len(toCheck)-1]
			res[idx] = i - idx
			toCheck = toCheck[:len(toCheck)-1]
		}
		toCheck = append(toCheck, i)
	}
	return res
}

// @lc code=end
func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
