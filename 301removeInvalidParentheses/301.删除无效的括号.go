/*
 * @lc app=leetcode.cn id=301 lang=golang
 *
 * [301] 删除无效的括号
 */
package main

import "fmt"

// @lc code=start
// 判断字符串是否是合法的括号字符串
func valid(v string) bool {
	cnt := 0
	for i := 0; i < len(v); i++ {
		if v[i] == '(' {
			cnt++
		} else if v[i] == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt <= 0
}

func removeInvalidParentheses(s string) []string {
	result := []string{}
	// bfs 使用map来避免重复的字符串
	var bfs func(queue map[string]bool) = func(queue map[string]bool) {
		// 队列不为空 则继续搜索
		for len(queue) != 0 {
			// 判断队列中的字符串是否为合法括号
			for k := range queue {
				if valid(k) {
					result = append(result, k)
				}
			}
			// 存在合法括号直接退出循环
			if len(result) != 0 {
				return
			}
			// 用来保存遍历queue 同时删除一个字符得到的字符串
			tempQueue := make(map[string]bool)
			// 遍历queue
			for k := range queue {
				bytes := []byte(k)
				// 每次删除一个字符 然后保存到tempQueue中
				for j := 0; j < len(k); j++ {
					temp := make([]byte, len(bytes))
					copy(temp, bytes)
					// 遍历所有情况
					if temp[j] == '(' || temp[j] == ')' {
						t := append(temp[:j], temp[j+1:]...)
						tempQueue[string(t)] = true
					}
				}
			}
			// 赋值给queue
			queue = tempQueue
		}
	}
	bfs(map[string]bool{s: true})
	return result
}

// @lc code=end
func main() {
	fmt.Println(removeInvalidParentheses("()())"))
}
