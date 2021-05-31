/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */
package main

import "fmt"

// 不妨假设整个数组的众数记做a，则最初的数组中a的数量大于其余所有数。当采用count计数的时候有两种情况：

// 1）假设candidate等于a，则当count从1变为0的过程，此区间内a的数量等于其余数的数量，因此以count=0为分界线，数组右端部分的众数仍然为a

// 2）假设candidate不等于a，则当count从1变为0的过程， 此区间内a的数量小于等于其余数的数量，因此以count=0为分界线，数组右端部分的众数仍然为a

// 因此，以count=0可以将整个原始数组分为若干部分，count=0右端部分的数组中的众数永远是a，最终必然会筛选出a
// @lc code=start
func majorityElement(nums []int) int {
	var candidate, count int
	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
		if v == candidate {
			count += 1
		} else {
			count -= 1
		}
	}
	return candidate
}

// @lc code=end
func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}
