package main

import (
	"fmt"
)

/**
剑指 Offer 06. 从尾到头打印链表
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。



示例 1：

输入：head = [1,3,2]
输出：[2,3,1]


限制：

0 <= 链表长度 <= 10000

思路：
1.递归
2.迭代
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

//递归
func reversePrint1(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	return append(reversePrint1(head.Next), head.Val)
}

//迭代
func reversePrint2(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	fmt.Println(reversePrint1(node1))
	fmt.Println(reversePrint2(node1))
}
