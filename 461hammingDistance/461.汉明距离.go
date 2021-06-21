/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */
package main

import "fmt"

// @lc code=start
func hammingDistance(x int, y int) int {
	res := x ^ y
	ans := 0
	for ; res != 0; res &= res - 1 {
		ans++
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(hammingDistance(3, 1))
}
