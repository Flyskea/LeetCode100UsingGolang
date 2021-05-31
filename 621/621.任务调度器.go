/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */
package main

import (
	"fmt"
	"sort"
)

/*
func leastInterval(tasks []byte, n int) (minTime int) {
	// 同一任务有多少个
    cnt := map[byte]int{}
    for _, t := range tasks {
        cnt[t]++
    }
	// 下一个可以执行的任务
    nextValid := make([]int, 0, len(cnt))
	// 该任务还剩的次数
    rest := make([]int, 0, len(cnt))
    for _, c := range cnt {
        nextValid = append(nextValid, 1)
        rest = append(rest, c)
    }

    for range tasks {
        minTime++
        minNextValid := math.MaxInt64
        for i, r := range rest {
            if r > 0 && nextValid[i] < minNextValid {
                minNextValid = nextValid[i]
            }
        }
        if minNextValid > minTime {
            minTime = minNextValid
        }
        best := -1
        for i, r := range rest {
            if r > 0 && nextValid[i] <= minTime && (best == -1 || r > rest[best]) {
                best = i
            }
        }
        nextValid[best] = minTime + n + 1
        rest[best]--
    }
    return
}
*/

// 例如：tasks = ["A","A","A","B","B","B"], n = 2
// 我们先安排出现次数最多的任务"A",并且让两次执行"A"的时间间隔为2。
// 在这个时间间隔内，我们用其他任务类型去填充，又因为其他任务类型只有"B"一个，
// 不够填充2的时间间隔，因此额外需要一个冷却时间间隔。
// 需要有n+1个不同任务才能填充冷却时间
// @lc code=start
func leastInterval(tasks []byte, n int) int {
	buckets := make([]int, 26)
	for i := range tasks {
		buckets[tasks[i]-'A']++
	}
	sort.Ints(buckets)
	maxTimes := buckets[25]
	maxCount := 1
	for i := 25; i >= 1; i-- {
		if buckets[i] == buckets[i-1] {
			maxCount++
		} else {
			break
		}
	}
	ans := (maxTimes-1)*(n+1) + maxCount
	if ans < len(tasks) {
		return len(tasks)
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
}
