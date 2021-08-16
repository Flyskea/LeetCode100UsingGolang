package main

import "fmt"

/* func leftChild(i int) int {
	return 2*i + 1
}

func percDown(nums []int, i, N int) {
	var (
		child, tmp int
	)
	for tmp = nums[i]; leftChild(i) < N; i = child {
		child = leftChild(i)
		if child != N-1 && nums[child+1] > nums[child] {
			child++
		}
		if tmp < nums[child] {
			nums[i] = nums[child]
		} else {
			break
		}
	}
	nums[i] = tmp
}

func findKthLargest(nums []int, k int) int {
	k = len(nums) - k + 1
	for i := len(nums) / 2; i >= 0; i-- {
		percDown(nums, i, len(nums))
	}
	for i := len(nums) - 1; i >= k; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		percDown(nums, 0, i)
	}
	return nums[0]
} */

func quickSelect(nums []int, k, left, right int) int {
	i, j := left+1, right
	if left == right {
		return nums[left]
	}
	pivot := nums[left]
	for {
		for i < right && nums[i] <= pivot {
			i++
		}
		for j > left && nums[j] >= pivot {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			break
		}
	}
	nums[left], nums[j] = nums[j], nums[left]
	if k == j {
		return nums[j]
	}
	if k > j {
		return quickSelect(nums, k, j+1, right)
	}
	return quickSelect(nums, k, left, j-1)
}

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, len(nums)-k, 0, len(nums)-1)
}

func main() {
	fmt.Println(findKthLargest([]int{1, 2, 6, 8, 4}, 3))
}
