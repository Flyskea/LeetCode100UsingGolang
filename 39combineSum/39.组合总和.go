/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */
package main

import (
	"fmt"
	"sort"
)

// @lc code=start
func backtracking(candidates, total []int, target, index int, ans *[][]int) {
	if target <= 0 {
		if target == 0 {
			*ans = append(*ans, append([]int{}, total...))
		}
	}
	for i := index; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		total = append(total, candidates[i])
		backtracking(candidates, total, target-candidates[i], i, ans)
		total = total[:len(total)-1]
	}
}
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	backtracking(candidates, []int{}, target, 0, &res)
	return res
}

// @lc code=end

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
