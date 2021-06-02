/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */
package main

import (
	"fmt"
	"strconv"
)

// @lc code=start
func decodeString(s string) string {
	stack, res, muti := []string{}, "", 0
	for _, v := range s {
		if v == '[' {
			stack = append(stack, fmt.Sprint(muti), res)
			res, muti = "", 0
		} else if v == ']' {
			curMuti, lastRes := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			tmp := ""
			c, _ := strconv.Atoi(curMuti)
			for i := 0; i < c; i++ {
				tmp += res
			}
			res = lastRes + tmp
		} else if v >= '0' && v <= '9' {
			muti = muti*10 + (int)(v-'0')
		} else {
			res += string(v)
		}
	}
	return res
}

// @lc code=end
func main() {
	fmt.Println(decodeString("3[a2[c]]"))
}
