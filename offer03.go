package main

import "fmt"

/**
找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3


限制：

2 <= n <= 100000

通过次数279,191提交次数413,656

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

思路：
方法一：哈希表 / Set
利用数据结构特点，容易想到使用哈希表（Set）记录数组的各个数字，当查找到重复数字则直接返回。

算法流程：
初始化： 新建 HashSet ，记为 dicdic ；
遍历数组 numsnums 中的每个数字 numnum ：
当 numnum 在 dicdic 中，说明重复，直接返回 numnum ；
将 numnum 添加至 dicdic 中；
返回 -1−1 。本题中一定有重复数字，因此这里返回多少都可以。
复杂度分析：
时间复杂度 O(N) ： 遍历数组使用 O(N) ，HashSet 添加与查找元素皆为 O(1) 。
空间复杂度 O(N) ： HashSet 占用 O(N) 大小的额外空间。

方法二：原地交换
题目说明尚未被充分使用，即 在一个长度为 n 的数组 nums 里的所有数字都在 0 ~ n-1 的范围内 。
此说明含义：数组元素的 索引 和 值 是 一对多 的关系。
因此，可遍历数组并通过交换操作，使元素的 索引 与 值 一一对应（即 nums[i] = inums[i]=i ）。
因而，就能通过索引映射对应的值，起到与字典等价的作用。

复杂度分析：
时间复杂度 O(N)： 遍历数组使用 O(N) ，每轮遍历的判断和交换操作使用 O(1) 。
空间复杂度 O(1) ： 使用常数复杂度的额外空间。
。
*/

func FindRepeatNum(nums []int) int {
	i := 0
	for i < len(nums) {
		fmt.Println("begin", nums)
		if nums[i] == i {
			fmt.Println("continue", nums)
			i++
			continue
		}
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		fmt.Println("end", nums)
	}
	return -1
}

func main() {
	fmt.Println(FindRepeatNum([]int{2, 3, 1, 0, 2, 5, 3}))
}
