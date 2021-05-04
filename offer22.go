package main

import (
	"fmt"
)

/**
剑指 Offer 22. 链表中倒数第k个节点
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。

例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。



示例：

给定一个链表: 1->2->3->4->5, 和 k = 2.

返回链表 4->5.

思路：
1.第一时间想到的解法：
先遍历统计链表长度，记为 nn ；
设置一个指针走 (n-k)步，即可找到链表倒数第 k个节点。
2.使用双指针则可以不用统计链表长度。
快指针先走k步，则当快指针走到终点，慢指针刚好走的n-k步，即倒数k个节点处

*/

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for ; k > 0; k-- {
		fast = fast.Next
	}
	for fast != nil {
		fast, slow = fast.Next, slow.Next
	}
	return slow
}

func main() {
	fmt.Println(printNumbers(3))
}
