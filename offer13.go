package main

import (
	"fmt"
)

/**
剑指 Offer 13. 机器人的运动范围
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，
它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。
例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。
请问该机器人能够到达多少个格子？

示例 1：

输入：m = 2, n = 3, k = 1
输出：3
示例 2：

输入：m = 3, n = 1, k = 0
输出：1
提示：

1 <= n,m <= 100
0 <= k <= 20

思路：
思路很简单:

访问过的点不可达，返回0
越界的点不可达，返回0
数位和大于k的点不可达，返回0
其余可达的点返回1，最后累加起来

*/

func getSum(m, n int) int {
	sum := 0
	for m != 0 {
		sum += m % 10
		m = m / 10
	}
	for n != 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}

func dfs(x, y, m, n, k int, vis [][]bool) int {
	if getSum(x, y) > k {
		return 0
	}
	if x >= m || y >= n || x < 0 || y < 0 {
		return 0
	}
	if vis[x][y] {
		return 0
	}
	vis[x][y] = true
	sum := 1
	sum += dfs(x-1, y, m, n, k, vis)
	sum += dfs(x+1, y, m, n, k, vis)
	sum += dfs(x, y-1, m, n, k, vis)
	sum += dfs(x, y+1, m, n, k, vis)
	return sum
}

func movingCount(m int, n int, k int) int {
	vis := make([][]bool, m+1)
	for i := 0; i < len(vis); i++ {
		vis[i] = make([]bool, n+1)
	}
	return dfs(0, 0, m, n, k, vis)
}
func main() {
	fmt.Println(movingCount(2, 3, 1))
	fmt.Println(movingCount(3, 1, 0))
}
