/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */
package main

import (
	"fmt"
)

// @lc code=start
func findAnagrams(s string, p string) []int {
	var freq [256]int
	result := []int{}
	if len(s) == 0 || len(s) < len(p) {
		return result
	}
	for i := 0; i < len(p); i++ {
		freq[p[i]-'a']++
	}
	left, right, count := 0, 0, len(p)

	for right < len(s) {
		if freq[s[right]-'a'] >= 1 {
			count--
		}
		freq[s[right]-'a']--
		right++
		if count == 0 {
			result = append(result, left)
		}
		if right-left == len(p) {
			if freq[s[left]-'a'] >= 0 {
				count++
			}
			freq[s[left]-'a']++
			left++
		}

	}
	return result
}

// @lc code=end
func main() {
	fmt.Println(findAnagrams("cbarbabacd", "abc"))
}
