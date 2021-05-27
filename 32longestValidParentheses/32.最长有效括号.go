/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */
package main

import (
	"fmt"
)

/*
栈
func longestValidParentheses(s string) int {
	stack := []int{}
	maxLen := 0
	// 方便求差
	stack = append(stack, -1)
	for i, v := range s {
		if v == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxLen = max(maxLen, i-stack[len(stack)-1])
			}
		}
	}
	return maxLen
} */

/*
动态规划
func longestValidParentheses(s string) int {
	sLen := len(s)
	dp := make([]int, sLen+1)
	maxLen := 0
	for i := 1; i < sLen; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				// 除去当前括号，将前面连续匹配的括号加上
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
				// i - dp[i-1]-1 最后一次匹配到括号的前一个
				// 如果是'('说明当前括号匹配上
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		maxLen = max(maxLen, dp[i])
	}
	return maxLen
} */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestValidParentheses(s string) int {
	left, right, maxLength := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLength
}

// @lc code=end

func main() {
	fmt.Println(longestValidParentheses("()(())"))
}
