package main

import "fmt"

/**
剑指 Offer 32 - I. 从上到下打印二叉树
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。



例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回：

[3,9,20,15,7]


提示：

节点总数 <= 1000

思路：
题目要求的二叉树的 从上至下 打印（即按层打印），又称为二叉树的 广度优先搜索（BFS）。
BFS 通常借助 队列 的先入先出特性来实现
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	var (
		queue []*TreeNode
		res   []int
	)
	queue = append(queue, root)
	for len(queue) > 0 {
		cur := queue[0]
		res = append(res, cur.Val)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
		queue = queue[1:]
	}
	return res
}

func main() {
	node := &TreeNode{Val: 1}
	node.Left = &TreeNode{Val: 2}
	node.Right = &TreeNode{Val: 3}
	node.Left.Left = &TreeNode{Val: 4}
	node.Right.Left = &TreeNode{Val: 5}
	node.Right.Right = &TreeNode{Val: 7}
	fmt.Println(levelOrder(node))
}
