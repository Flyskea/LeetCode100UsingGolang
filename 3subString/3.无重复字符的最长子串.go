/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
package main

import "fmt"

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func lengthOfLongestSubstring(s string) int {
	sLen := len(s)
	if sLen == 0 {
		return 0
	}
	var freq [256]byte
	result, left, right := 0, 0, -1

	for left < sLen {
		if right+1 < sLen && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++
		} else {
			freq[s[left]]--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

// @lc code=end

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
