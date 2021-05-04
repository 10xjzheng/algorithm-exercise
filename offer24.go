package main

import (
	"fmt"
)

/**
剑指 Offer 24. 反转链表
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。



示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL


限制：

0 <= 节点个数 <= 5000

思路：
需要pre，cur，next 3个辅助指针，循环里面需要做两个步骤：
1.存下cur->next为next，并将cur->next 指向 pre
2.将pre变成cur，将cur变成next
直到cur为nil则退出，取pre
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	pre, cur := nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre, cur = cur, next
	}
	return pre
}

//递归
func reverseListByRecur(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func main() {
	node := &ListNode{Val: 1}
	node.Next = &ListNode{Val: 2}
	node.Next.Next = &ListNode{Val: 3}
	fmt.Println(reverseList(node))
}
