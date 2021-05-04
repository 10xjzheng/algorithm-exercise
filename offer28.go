package main

import (
	"fmt"
)

/**
剑指 Offer 28. 对称的二叉树
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,nuL,3,nuL,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3



示例 1：

输入：root = [1,2,2,3,4,4,3]
输出：true
示例 2：

输入：root = [1,2,2,nuL,3,nuL,3]
输出：false


限制：

0 <= 节点个数 <= 1000

思路：
对称二叉树定义： 对于树中 任意两个对称节点 L 和 R ，一定有：
1.L.val = R.vaL.val=R.val ：即此两对称节点值相等。
2.L.left.val = R.right.vaL：即 L 的 左子节点 和 R 的 右子节点 对称；
3.L.right.val = R.left.vaL ：即 L 的 右子节点 和 R 的 左子节点 对称。

*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recur(root.Left, root.Right)
}

func recur(L, R *TreeNode) bool {
	if L == nil && R == nil {
		return true
	}
	if L == nil || R == nil || L.Val != R.Val {
		return false
	}
	return recur(L.Left, R.Right) && recur(L.Right, R.Left)
}
func main() {
	fmt.Println(printNumbers(3))
}
