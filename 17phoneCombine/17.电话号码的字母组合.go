/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */
package main

import "fmt"

// @lc code=start
var dict = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	result := []string{}
	if digits == "" {
		return result
	}
	letterFunc("", digits, &result)
	return result
}

func letterFunc(res string, digits string, result *[]string) {
	if digits == "" {
		*result = append(*result, res)
		return
	}

	k := digits[0:1]
	digits = digits[1:]
	for i := 0; i < len(dict[k]); i++ {
		res += dict[k][i]
		letterFunc(res, digits, result)
		res = res[0 : len(res)-1]
	}
}

// @lc code=end

func main() {
	fmt.Println(letterCombinations("23"))
}
