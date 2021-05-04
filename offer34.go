package main

import "fmt"

/**
剑指 Offer 34. 二叉树中和为某一值的路径
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。



示例:
给定如下二叉树，以及目标和 target = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]


提示：

节点总数 <= 10000

思路：
本问题是典型的二叉树方案搜索问题，使用回溯法解决，其包含 先序遍历 + 路径记录 两部分。

先序遍历： 按照 “根、左、右” 的顺序，遍历树的所有节点。
路径记录： 在先序遍历中，记录从根节点到当前节点的路径。当路径为 ① 根节点到叶节点形成的路径 且 ② 各节点值的和等于目标值 sum 时，将此路径加入结果列表。
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
func pathSum(root *TreeNode, target int) [][]int {
	res := make([][]int, 0, 0)
	if root == nil {
		return res
	}

	return dfs(root, target, []int{}, res)
}

// root 当前节点
// target 当前值
// path   当前走的路径
// lastRes 当前路径结果集
// return: 遍历后的结果集
func dfs(root *TreeNode, target int, path []int, lastRes [][]int) [][]int {
	if root == nil {
		return lastRes
	}
	// 计算新的值
	target = target - root.Val
	// 添加路径
	path = append(path, root.Val)
	// 如果找到了路径，并且是根节点就添加进结果集
	if target == 0 && root.Left == nil && root.Right == nil {
		// var tempPath []int
		// tempPath = append(tempPath, path...)
		// lastRes = append(lastRes, tempPath)
		// 添加path到结果集
		lastRes = append(lastRes, append([]int{}, path...))
	}

	// 先序遍历 依次为 根-> 左-> 右
	lastRes = dfs(root.Left, target, path, lastRes)
	lastRes = dfs(root.Right, target, path, lastRes)

	// 移除遍历后的路径
	path = path[:len(path)-1]

	return lastRes
}
func main() {
}
