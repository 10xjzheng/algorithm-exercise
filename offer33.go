package main

import "fmt"

/**
剑指 Offer 33. 二叉搜索树的后序遍历序列
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。



参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [1,6,3,2,5]
输出: false
示例 2：

输入: [1,3,2,6,5]
输出: true

思路：
1.如果切片<=1,肯定是true
2.先找到每棵树的根节点 root = postorder[len(postorder) - 1]，同时找到第一个大于root的值index作为左右子树的分割点
3.确定左右子树：leftArr := postorder[:index]，rigth := postorder[index:len(postorder)-1]
4.不为二叉树的条件就是左子树的值大于根节点，右子树的小于根节点
5.递归判断左右子树

*/
func verifyPostorder(postorder []int) bool {
	length := len(postorder)
	if length <= 1 {
		return true
	}
	rootVal := postorder[length-1]
	i := 0
	for ; i < length; i++ {
		//这里的 >= 不能是 > ，不然会死循环。
		if postorder[i] >= rootVal {
			break
		}
	}
	r := i
	for ; i < length; i++ {
		if postorder[i] < rootVal {
			return false
		}
	}
	return verifyPostorder(postorder[:r]) && verifyPostorder(postorder[r:length-1])
}
func main() {
	fmt.Println(verifyPostorder([]int{1, 2, 5, 10, 6, 9, 4, 3}))
}
