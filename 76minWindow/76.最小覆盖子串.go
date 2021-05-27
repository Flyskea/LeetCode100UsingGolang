/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */
package main

import (
	"fmt"
)

// @lc code=start
func minWindow(s string, t string) string {
	chars := make([]int, 128)
	flag := make([]bool, 128)
	for i := range t {
		flag[t[i]] = true
		chars[t[i]]++
	}
	var cnt, l, minL, minSize int
	minSize = len(s) + 1
	for r := 0; r < len(s); r++ {
		if flag[s[r]] {
			chars[s[r]]--
			if chars[s[r]] >= 0 {
				cnt++
			}
			// 若目前滑动窗口已包含T中全部字符，
			// 则尝试将l右移，在不影响结果的情况下获得最短子字符串
			for cnt == len(t) {
				if r-l+1 < minSize {
					minL = l
					minSize = r - l + 1
				}
				if flag[s[l]] {
					chars[s[l]]++
					if chars[s[l]] > 0 {
						cnt--
					}
				}
				l++
			}
		}
	}
	if minSize > len(s) {
		return ""
	}
	return s[minL : minL+minSize]
}

// @lc code=end
func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
