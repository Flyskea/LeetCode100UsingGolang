/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 */
package main

import "fmt"

/*
// DP
// func countSubstrings(s string) int {
	sLen := len(s)
	dp := make([][]bool, sLen)
	ans := 0
	for i := range dp {
		dp[i] = make([]bool, sLen)
	}
	for i := sLen - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j]) && ((j-i) < 3 || dp[i+1][j-1])
			if dp[i][j] {
				ans++
			}
		}
	}
	return ans
}
*/
// @lc code=start
func countSubstrings(s string) int {
	ans := 0
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		for j := 0; j < 2; j++ {
			left := i
			right := i + j
			for left >= 0 && right < sLen && s[left] == s[right] {
				ans++
				left--
				right++
			}
		}
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(countSubstrings("aaa"))
}
