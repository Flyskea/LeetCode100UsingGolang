/*
 * @lc app=leetcode.cn id=301 lang=golang
 *
 * [301] 删除无效的括号
 */
package main

import "fmt"

// @lc code=start
/* // 判断字符串是否是合法的括号字符串
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

func bfs(result *[]string, queue map[string]bool) {
	// 队列不为空 则继续搜索
	for len(queue) != 0 {
		// 判断队列中的字符串是否为合法括号
		for k := range queue {
			if valid(k) {
				*result = append(*result, k)
			}
		}
		// 存在合法括号直接退出循环
		if len(*result) != 0 {
			return
		}
		// 用来保存遍历queue 同时删除一个字符得到的字符串
		tempQueue := make(map[string]bool)
		// 遍历queue
		for k := range queue {
			// 每次删除一个字符 然后保存到tempQueue中
			for j := 0; j < len(k); j++ {
				// 遍历所有情况
				if k[j] == '(' || k[j] == ')' {
					t := append([]byte{}, []byte(k[:j])...)
					t = append(t, []byte(k[j+1:])...)
					tempQueue[string(t)] = true
				}
			}
		}
		// 赋值给queue
		queue = tempQueue
	}
}

func removeInvalidParentheses(s string) []string {
	result := []string{}
	// bfs 使用map来避免重复的字符串
	bfs(&result, map[string]bool{s: true})
	return result
} */

var m map[string]bool

func removeInvalidParentheses(s string) []string {
	if len(s) == 0 {
		return []string{""}
	}
	leftRemove := 0
	rightRemove := 0
	for _, v := range s {
		if v == '(' {
			leftRemove++
		} else if v == ')' {
			//fixme：下面这个参数很关键【参考自答案】，如：))((【这时leftRemove=2，rightRemove=2】
			if leftRemove == 0 {
				rightRemove++
			} else {
				leftRemove--
			}
		}
	}

	m = make(map[string]bool)
	dfs(s, 0, len(s), 0, 0, leftRemove, rightRemove, "")

	ret := make([]string, 0)
	for v := range m {
		ret = append(ret, v)
	}
	return ret
}

func dfs(s string, index, length, leftCount, rightCount, leftRemove, rightRemove int, path string) {
	if index == length {
		if leftRemove == 0 && rightRemove == 0 {
			m[path] = true
		}
		return
	}
	c := s[index]
	//todo: 参考答案的示例，最好自己画下图，深刻理解下！！！【通过index++来最终跳出递归！！！】
	//1、去掉左括号
	if c == '(' && leftRemove > 0 {
		dfs(s, index+1, length, leftCount, rightCount, leftRemove-1, rightRemove, path)
	}
	//2、去掉右括号
	if c == ')' && rightRemove > 0 {
		dfs(s, index+1, length, leftCount, rightCount, leftRemove, rightRemove-1, path)
	}
	//3、保留字符
	path += string(c)
	if c != '(' && c != ')' {
		dfs(s, index+1, length, leftCount, rightCount, leftRemove, rightRemove, path)
	} else if c == '(' {
		dfs(s, index+1, length, leftCount+1, rightCount, leftRemove, rightRemove, path)
	} else if c == ')' && leftCount > rightCount {
		//增加右括号时，必须加上 leftCount > rightCount 这个判断，不然左括号数量少于右括号是不合法的！！！！
		dfs(s, index+1, length, leftCount, rightCount+1, leftRemove, rightRemove, path)
	}
}

// @lc code=end
func main() {
	fmt.Println(removeInvalidParentheses("()())"))
}
