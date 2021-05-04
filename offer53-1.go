package main

import (
	"fmt"
)

/**
剑指 Offer 53 - I. 在排序数组中查找数字 I
统计一个数字在排序数组中出现的次数。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: 2
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: 0


限制：

0 <= 数组长度 <= 50000

思路：
排序数组 nums 中的所有数字 target 形成一个窗口，记窗口的 左 / 右边界 索引分别为 left和 right ，
分别对应窗口左边 / 右边的首个元素。
本题要求统计数字target 的出现次数，可转化为：使用二分法分别找到 左边界 left和 和 右边界 right ，
易得数字 target 的数量为right−left−1 。
*/
func search(nums []int, target int) int {
	fmt.Println(findIndex(nums, target))
	fmt.Println(findIndex(nums, target-1))
	return findIndex(nums, target) - findIndex(nums, target-1)
}

func findIndex(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] <= target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}
func main() {
	fmt.Println(search([]int{1, 2, 3, 6, 6, 6, 8, 9, 11}, 6))
}
