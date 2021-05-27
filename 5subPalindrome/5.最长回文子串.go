/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
package main

import "fmt"

/* 动态规划
根据这样的思路，我们就可以用动态规划的方法解决本题。我们用P(i,j)表示字符串s的第i到j个字母组成的串
{
	P(i,i)=true
	// 当子串长度为2时
	P(i,i+1)=(Si==Si+1)
}
// 最优公式
P(i,j)=P(i+1,j−1)∧(Si==Sj)
func longestPalindrome(s string) string {
	sLen := len(s)
	res, dp := "", make([][]bool, sLen)
	for i := range dp {
		dp[i] = make([]bool, sLen)
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			// 子串长度小于3直接判断
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
} */

// @lc code=start
// 我们可以从每一种边界情况开始「扩展」，也可以得出所有的状态对应的答案。
// 边界情况即为子串长度为 11 或 22 的情况。
// 我们枚举每一种边界情况，并从对应的子串开始不断地向两边扩展。
// 如果两边的字母相同，我们就可以继续扩展，例如从 P(i+1,j-1)P(i+1,j−1) 扩展到 P(i,j)P(i,j)；
// 如果两边的字母不同，我们就可以停止扩展，因为在这之后的子串都不能是回文串了。
// 中心拓展
func maxPalindrome(s string, i, j int, res string) string {
	sub := ""
	for j < len(s) && i >= 0 && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(res) < len(sub) {
		return sub
	}
	return res
}

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res = maxPalindrome(s, i, i, res)
		res = maxPalindrome(s, i, i+1, res)
	}
	return res
}

// @lc code=end

func main() {
	fmt.Println(longestPalindrome("babad"))
}
