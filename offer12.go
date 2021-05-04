package main

import "fmt"

/**
剑指 Offer 12. 矩阵中的路径
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。

示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：

输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false


提示：

1 <= board.length <= 200
1 <= board[i].length <= 200
board 和 word 仅由大小写英文字母组成

思路：
本问题是典型的矩阵搜索问题，可使用 深度优先搜索（DFS）+ 剪枝 解决。

深度优先搜索： 可以理解为暴力法遍历矩阵中所有字符串可能性。DFS 通过递归，先朝一个方向搜到底，再回溯至上个节点，沿另一个方向搜索，以此类推。
剪枝： 在搜索中，遇到 这条路不可能和目标字符串匹配成功 的情况（例如：此矩阵元素和目标字符不同、此元素已被访问），则应立即返回，称之为 可行性剪枝 。

DFS 解析：
递归参数： 当前元素在矩阵 board 中的行列索引 i 和 j ，当前目标字符在 word 中的索引 k 。
终止条件：
返回 false ： (1) 行或列索引越界 或 (2) 当前矩阵元素与目标字符不同 或 (3) 当前矩阵元素已访问过 （ (3) 可合并至 (2) ） 。
返回 true ： k = len(word) - 1 ，即字符串 word 已全部匹配。
递推工作：
标记当前矩阵元素： 将 board[i][j] 修改为 空字符 '' ，代表此元素已访问过，防止之后搜索时重复访问。
搜索下一单元格： 朝当前元素的 上、下、左、右 四个方向开启下层递归，使用 或 连接 （代表只需找到一条可行路径就直接返回，不再做后续 DFS ），并记录结果至 res 。
还原当前矩阵元素： 将 board[i][j] 元素还原至初始值，即 word[k] 。
返回值： 返回布尔量 res ，代表是否搜索到目标字符串。

*/
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//如果在数组中找得到第一个数，就执行下一步，否则返回false
			if search(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}
func search(board [][]byte, i, j, k int, word string) bool {
	//如果找到最后一个数，则返回true,搜索成功
	if k == len(word) {
		return true
	}
	//i,j的约束条件
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return false
	}

	//进入DFS深度优先搜索
	//先把正在遍历的该值重新赋值，如果在该值的周围都搜索不到目标字符，则再把该值还原
	//如果在数组中找到第一个字符，则进入下一个字符的查找
	if board[i][j] == word[k] {
		temp := board[i][j]
		board[i][j] = ' '
		//下面这个if语句，如果成功进入，说明找到该字符，然后进行下一个字符的搜索,直到所有的搜索都成功，
		//即k == len(word) - 1 的大小时，会返回true，进入该条件语句，然后返回函数true值。
		if search(board, i, j+1, k+1, word) ||
			search(board, i, j-1, k+1, word) ||
			search(board, i+1, j, k+1, word) ||
			search(board, i-1, j, k+1, word) {
			return true
		} else {
			//没有找到目标字符，还原之前重新赋值的board[i][j]
			board[i][j] = temp
		}
	}
	return false
}

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Println(exist(board, "ABCCED"))
	fmt.Println(exist(board, "DFCED"))
}
