package main

import (
	"fmt"
)

/**
剑指 Offer 25. 合并两个排序的链表
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
限制：

0 <= 链表长度 <= 1000

思路：
1.递归
2.迭代
*/
func mergeTwoListsByRecur(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoListsByRecur(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoListsByRecur(l1, l2.Next)
		return l2
	}
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var head *ListNode
	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next, l1 = l1, l1.Next
		} else {
			cur.Next, l2 = l2, l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head
}

func main() {
	fmt.Println(printNumbers(3))
}
