package main

import (
	"fmt"
)

/**
剑指 Offer 30. 包含min函数的栈
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。



示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.


提示：

各函数的调用总次数不超过 20000 次

思路：
因为原栈stack和min的最小值下标位置是对应的，所以将最小值放在min最后，最后直接返回min的最后一个值即可
*/
type MinStack struct {
	stack []int
	min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min:   []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.min) == 0 || x <= this.min[len(this.min)-1] { // 没有元素或者x小于最小值的时候，append到min最后
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, this.min[len(this.min)-1]) // 否则继续将上一个最小值append到min
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	if len(this.min) > 0 {
		return this.min[len(this.min)-1]
	}
	return 0
}

func main() {
	fmt.Println(printNumbers(3))
}
