package main

import (
	"fmt"
)

/**
剑指 Offer 07. 重建二叉树
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7


限制：

0 <= 节点个数 <= 5000

思路：
二叉树前序遍历的顺序为：

先遍历根节点；

随后递归地遍历左子树；

最后递归地遍历右子树。

二叉树中序遍历的顺序为：

先递归地遍历左子树；

随后遍历根节点；

最后递归地遍历右子树。

在「递归」地遍历某个子树的过程中，我们也是将这颗子树看成一颗全新的树，按照上述的顺序进行遍历。挖掘「前序遍历」和「中序遍历」的性质，我们就可以得出本题的做法。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

func main() {
	preOrder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preOrder, inorder)
	fmt.Println(root, root.Left, root.Right, root.Right.Left, root.Right.Right)
}
