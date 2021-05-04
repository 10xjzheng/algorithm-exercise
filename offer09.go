package main

import (
	"container/list"
	"fmt"
)

/**
剑指 Offer 09. 用两个栈实现队列
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )



示例 1：

输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]
示例 2：

输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
提示：

1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用
*/

type Queue struct {
	stack1, stack2 *list.List
}

func NewQueue() *Queue {
	return &Queue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (q *Queue) DeQueue() int {
	if q.stack2.Len() == 0 {
		for q.stack1.Len() > 0 {
			q.stack2.PushBack(q.stack1.Remove(q.stack1.Back()))
		}
	}
	if q.stack2.Len() != 0 {
		e := q.stack2.Back()
		q.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}

func (q *Queue) InQueue(num int) {
	q.stack1.PushBack(num)
}

func main() {
	q := NewQueue()
	q.InQueue(1)
	q.InQueue(2)
	q.InQueue(3)
	q.InQueue(4)
	fmt.Println(q.DeQueue())
	q.InQueue(5)
	fmt.Println(q.DeQueue())
	q.InQueue(6)
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
}
