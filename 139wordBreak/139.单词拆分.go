/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */
package main

import "fmt"

func DFS(start int, s string, wordMap map[string]bool) bool {
	sLen := len(s)
	if start == sLen {
		return true
	}
	for i := start + 1; i <= sLen; i++ {
		prefix := s[start:i]
		if wordMap[prefix] && DFS(i, s, wordMap) {
			return true
		}
	}
	return false
}
func wordBreakDFS(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, v := range wordDict {
		wordMap[v] = true
	}
	return DFS(0, s, wordMap)
}

func wordBreakBFS(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, v := range wordDict {
		wordMap[v] = true
	}
	queue := []int{0}
	sLen := len(s)
	start := 0
	prefix := ""
	for len(queue) != 0 {
		start = queue[0]
		queue = queue[1:]
		for i := start + 1; i <= sLen; i++ {
			prefix = s[start:i]
			if wordMap[prefix] {
				if i < sLen {
					queue = append(queue, i)
				} else {
					return true
				}
			}
		}
	}
	return false
}

// @lc code=start
// DP
func wordBreak(s string, wordDict []string) bool {
	sLen := len(s)
	dp := make([]bool, sLen+1)
	prefix := ""
	wordMap := make(map[string]bool)
	for _, v := range wordDict {
		wordMap[v] = true
	}
	dp[0] = true
	for i := 1; i <= sLen; i++ {
		for j := i - 1; j >= 0; j-- {
			if !dp[j] {
				continue
			}
			prefix = s[j:i]
			if wordMap[prefix] && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	ans := dp[sLen]
	return ans
}

// @lc code=end
func main() {
	fmt.Println(wordBreakBFS("leetcode", []string{"leet", "code"}))
}
