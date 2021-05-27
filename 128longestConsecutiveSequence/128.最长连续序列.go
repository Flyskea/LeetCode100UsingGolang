/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */
package main

import "fmt"

// @lc code=start
func longestConsecutive(nums []int) int {
	var currentNum, currentSeq, max int
	numSet := make(map[int]bool)
	for _, v := range nums {
		numSet[v] = true
	}
	for i := range numSet {
		if !numSet[i-1] {
			currentNum = i
			currentSeq = 1
			for numSet[currentNum+1] {
				currentNum++
				currentSeq++
			}
			if max < currentSeq {
				max = currentSeq
			}
		}
	}
	return max
}

// @lc code=end
func main() {
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
