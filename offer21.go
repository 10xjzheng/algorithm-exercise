package main

import (
	"fmt"
)

/**
剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。



示例：

输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。


提示：

0 <= nums.length <= 50000
1 <= nums[i] <= 10000

思路：
首尾双指针
定义头指针 left ，尾指针 right .
left 一直往右移，直到它指向的值为偶数
right 一直往左移， 直到它指向的值为奇数
交换 nums[left] 和 nums[right] .
重复上述操作，直到 left == right.

*/

func exchange(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l]&1 == 1 {
			l++
			continue
		}
		if nums[r]&1 == 0 {
			r--
			continue
		}
		tmp := nums[l]
		nums[l] = nums[r]
		nums[r] = tmp
	}
	return nums
}

func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4}))
}
