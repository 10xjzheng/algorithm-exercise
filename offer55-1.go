package main

import (
	"fmt"
)

/**
剑指 Offer 55 - I. 二叉树的深度
输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

例如：

给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。

思路：
如果二叉树为空，深度为 0；
如果二叉树只有根节点，深度为 1；
如果二叉树的根节点只有左子树，深度为左子树的深度加 1；
如果二叉树的根节点只有右子树，深度为右子树的深度加 1；
如果二叉树的根节点既有左子树又有右子树，深度为左右子树深度的最大者再加 1。

*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}
func main() {
	fmt.Println(printNumbers(3))
}
