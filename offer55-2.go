package main

import (
	"fmt"
	"math"
)

/**
剑指 Offer 55 - II. 平衡二叉树
输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。



示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。



限制：

0 <= 树的结点个数 <= 10000

思路：
 此树的深度 等于 左子树的深度 与 右子树的深度 中的 最大值 +1+1

方法一：后序遍历 + 剪枝 （从底至顶）
此方法为本题的最优解法，但剪枝的方法不易第一时间想到。

思路是对二叉树做后序遍历，从底至顶返回子树深度，若判定某子树不是平衡树则 “剪枝” ，直接向上返回。

算法流程：
recur(root) 函数：

返回值：
当节点root 左 / 右子树的深度差 \leq 1≤1 ：则返回当前子树的深度，即节点 root 的左 / 右子树的深度最大值 +1+1 （ max(left, right) + 1 ）；
当节点root 左 / 右子树的深度差 > 2>2 ：则返回 -1−1 ，代表 此子树不是平衡树 。
终止条件：
当 root 为空：说明越过叶节点，因此返回高度 00 ；
当左（右）子树深度为 -1−1 ：代表此树的 左（右）子树 不是平衡树，因此剪枝，直接返回 -1−1 ；
isBalanced(root) 函数：

返回值： 若 recur(root) != -1 ，则说明此树平衡，返回 truetrue ； 否则返回 falsefalse 。

作者：jyd
链接：https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/solution/mian-shi-ti-55-ii-ping-heng-er-cha-shu-cong-di-zhi/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func isBalanced(node *TreeNode) bool {
	return find(node) != -1
}

func find(node *TreeNode) float64 {
	if node == nil {
		return 0
	}

	l := find(node.Left) //!左节点
	if l == -1 {         //剪枝，不平衡时直接返回，
		return -1
	}

	r := find(node.Right) //!右节点
	if r == -1 {          //剪枝，不平衡时直接返回
		return -1
	}

	if math.Abs(l-r) > 1 { //剪枝，不平衡时直接返回
		return -1
	}

	return math.Max(l, r) + 1 //计算深度 !根节点
}

/*作者：bygo
链接：https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/solution/zi-ding-xiang-xia-zi-di-xiang-shang-by-linbingyu-2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。*/
func main() {
	fmt.Println(printNumbers(3))
}
