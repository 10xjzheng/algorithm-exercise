package main

/**
剑指 Offer 26. 树的子结构
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:
给定的树 A:

     3
    / \
   4   5
  / \
 1   2
给定的树 B：

   4
  /
 1
返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

示例 1：

输入：A = [1,2,3], B = [3,1]
输出：false
示例 2：

输入：A = [3,4,5,1,2], B = [4,1]
输出：true
限制：

0 <= 节点个数 <= 10000

思路：
若树 B 是树 A 的子结构，则子结构的根节点可能为树 A 的任意一个节点。因此，判断树 B 是否是树 A 的子结构，需完成以下两步工作：

1.先序遍历树 A 中的每个节点 n_An（对应函数 isSubStructure(A, B)）
2.判断树 A 中 以 n_An为根节点的子树 是否包含树 B 。（对应函数 recur(A, B)）
recur(A, B) 函数：

终止条件：
当节点 B 为空：说明树 B 已匹配完成（越过叶子节点），因此返回 true ；
当节点 A 为空：说明已经越过树 A 叶子节点，即匹配失败，返回 false ；
当节点 A 和 B 的值不同：说明匹配失败，返回 false ；
返回值：
判断 A 和 B 的左子节点是否相等，即 recur(A.left, B.left) ；
判断 A 和 B 的右子节点是否相等，即 recur(A.right, B.right) ；
*/
func isSubStructure(A, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

//逐一对比每个节点是否一致
func recur(A, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}
func main() {
}
