package main

import (
	"fmt"
)

/**
剑指 Offer 52. 两个链表的第一个公共节点

解题思路：
我们使用两个指针 node1，node2 分别指向两个链表 headA，headB 的头结点，然后同时分别逐结点遍历，
当 node1 到达链表 headA 的末尾时，重新定位到链表 headB 的头结点；当 node2 到达链表 headB 的末尾时，
重新定位到链表 headA 的头结点。

这样，当它们相遇时，所指向的结点就是第一个公共结点。

两个链表长度分别为L1+C、L2+C， C为公共部分的长度，按照楼主的做法： 第一个人走了L1+C步后，回到第二个人起点走L2步；
第2个人走了L2+C步后，回到第一个人起点走L1步。 当两个人走的步数都为L1+L2+C时就两个家伙就相爱了
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	node1, node2 := headA, headB
	for node1 != node2 {
		if node1 != nil {
			node1 = node1.Next
		} else {
			node1 = headB
		}
		if node2 != nil {
			node2 = node2.Next
		} else {
			node2 = headA
		}
	}
	return node1
}
func main() {
	fmt.Println(printNumbers(3))
}
