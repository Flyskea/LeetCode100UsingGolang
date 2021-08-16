package main

import "fmt"

func leftChild(i int) int {
	return 2*i + 1
}

func percDown(nums []int, i, N int, occurrences map[int]int) {
	var (
		child, tmp int
	)
	for tmp = nums[i]; leftChild(i) < N; i = child {
		child = leftChild(i)
		if child != N-1 && occurrences[nums[child+1]] < occurrences[nums[child]] {
			child++
		}
		if occurrences[tmp] > occurrences[nums[child]] {
			nums[i] = nums[child]
		} else {
			break
		}
	}
	nums[i] = tmp
}

func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	tmp := []int{}
	for k := range occurrences {
		tmp = append(tmp, k)
	}
	for i := len(tmp) / 2; i >= 0; i-- {
		percDown(tmp, i, len(tmp), occurrences)
	}
	for i := len(tmp) - 1; i >= k; i-- {
		tmp[0], tmp[i] = tmp[i], tmp[0]
		percDown(tmp, 0, i, occurrences)
	}
	ans := []int{}
	for i := 0; i < k; i++ {
		ans = append(ans, tmp[i])
	}
	return ans
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 1))
}
